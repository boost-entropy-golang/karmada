/*
Copyright 2022 The Karmada Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package proxy

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/diff"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/kubernetes/scheme"
	restclient "k8s.io/client-go/rest"

	clusterv1alpha1 "github.com/karmada-io/karmada/pkg/apis/cluster/v1alpha1"
	policyv1alpha1 "github.com/karmada-io/karmada/pkg/apis/policy/v1alpha1"
	searchv1alpha1 "github.com/karmada-io/karmada/pkg/apis/search/v1alpha1"
	karmadafake "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/fake"
	karmadainformers "github.com/karmada-io/karmada/pkg/generated/informers/externalversions"
	"github.com/karmada-io/karmada/pkg/search/proxy/framework"
	pluginruntime "github.com/karmada-io/karmada/pkg/search/proxy/framework/runtime"
	"github.com/karmada-io/karmada/pkg/search/proxy/store"
	proxytest "github.com/karmada-io/karmada/pkg/search/proxy/testing"
	"github.com/karmada-io/karmada/pkg/util"
)

func TestController(t *testing.T) {
	restConfig := &restclient.Config{
		Host: "https：//localhost:6443",
	}

	cluster1 := newCluster("cluster1")
	rr := &searchv1alpha1.ResourceRegistry{
		ObjectMeta: metav1.ObjectMeta{Name: "rr"},
		Spec: searchv1alpha1.ResourceRegistrySpec{
			ResourceSelectors: []searchv1alpha1.ResourceSelector{
				proxytest.PodSelector,
			},
		},
	}

	kubeFactory := informers.NewSharedInformerFactory(fake.NewSimpleClientset(), 0)
	karmadaFactory := karmadainformers.NewSharedInformerFactory(karmadafake.NewSimpleClientset(cluster1, rr), 0)

	ctrl, err := NewController(NewControllerOption{
		RestConfig:        restConfig,
		RestMapper:        proxytest.RestMapper,
		KubeFactory:       kubeFactory,
		KarmadaFactory:    karmadaFactory,
		MinRequestTimeout: 0,
	})

	if err != nil {
		t.Error(err)
		return
	}
	if ctrl == nil {
		t.Error("ctrl is nil")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	kubeFactory.Start(ctx.Done())
	karmadaFactory.Start(ctx.Done())
	ctrl.Start(ctx)
	defer ctrl.Stop()

	kubeFactory.WaitForCacheSync(ctx.Done())
	karmadaFactory.WaitForCacheSync(ctx.Done())
	// wait for controller synced
	time.Sleep(time.Second)

	hasPod := ctrl.store.HasResource(proxytest.PodGVR)
	if !hasPod {
		t.Error("has no pod resource")
		return
	}
}

func TestController_reconcile(t *testing.T) {
	newMultiNs := func(namespaces ...string) *store.MultiNamespace {
		multiNs := store.NewMultiNamespace()
		if len(namespaces) == 0 {
			multiNs.Add(metav1.NamespaceAll)
			return multiNs
		}
		for _, ns := range namespaces {
			multiNs.Add(ns)
		}
		return multiNs
	}
	tests := []struct {
		name  string
		input []runtime.Object
		want  map[string]map[string]*store.MultiNamespace
	}{
		{
			name:  "all empty",
			input: []runtime.Object{},
			want:  map[string]map[string]*store.MultiNamespace{},
		},
		{
			name: "resource registered, while cluster not registered",
			input: []runtime.Object{
				&searchv1alpha1.ResourceRegistry{
					ObjectMeta: metav1.ObjectMeta{Name: "rr1"},
					Spec: searchv1alpha1.ResourceRegistrySpec{
						TargetCluster: policyv1alpha1.ClusterAffinity{
							ClusterNames: []string{"cluster1"},
						},
						ResourceSelectors: []searchv1alpha1.ResourceSelector{
							proxytest.PodSelector,
						},
					},
				},
			},
			want: map[string]map[string]*store.MultiNamespace{},
		},
		{
			name: "pod and node are registered",
			input: []runtime.Object{
				newCluster("cluster1"),
				newCluster("cluster2"),
				&searchv1alpha1.ResourceRegistry{
					ObjectMeta: metav1.ObjectMeta{Name: "rr1"},
					Spec: searchv1alpha1.ResourceRegistrySpec{
						TargetCluster: policyv1alpha1.ClusterAffinity{
							ClusterNames: []string{"cluster1", "cluster2"},
						},
						ResourceSelectors: []searchv1alpha1.ResourceSelector{
							proxytest.PodSelector,
							proxytest.NodeSelector,
						},
					},
				},
			},
			want: map[string]map[string]*store.MultiNamespace{
				"cluster1": {
					"pods":  newMultiNs(),
					"nodes": newMultiNs(),
				},
				"cluster2": {
					"pods":  newMultiNs(),
					"nodes": newMultiNs(),
				},
			},
		},
		{
			name: "register pod in cluster1, register node in cluster2",
			input: []runtime.Object{
				newCluster("cluster1"),
				newCluster("cluster2"),
				&searchv1alpha1.ResourceRegistry{
					ObjectMeta: metav1.ObjectMeta{Name: "rr1"},
					Spec: searchv1alpha1.ResourceRegistrySpec{
						TargetCluster:     policyv1alpha1.ClusterAffinity{ClusterNames: []string{"cluster1"}},
						ResourceSelectors: []searchv1alpha1.ResourceSelector{proxytest.PodSelector},
					},
				},
				&searchv1alpha1.ResourceRegistry{
					ObjectMeta: metav1.ObjectMeta{Name: "rr2"},
					Spec: searchv1alpha1.ResourceRegistrySpec{
						TargetCluster:     policyv1alpha1.ClusterAffinity{ClusterNames: []string{"cluster2"}},
						ResourceSelectors: []searchv1alpha1.ResourceSelector{proxytest.NodeSelector},
					},
				},
			},
			want: map[string]map[string]*store.MultiNamespace{
				"cluster1": {
					"pods": newMultiNs(),
				},
				"cluster2": {
					"nodes": newMultiNs(),
				},
			},
		},
		{
			name: "register pod,node in cluster1, register node in cluster2",
			input: []runtime.Object{
				newCluster("cluster1"),
				newCluster("cluster2"),
				&searchv1alpha1.ResourceRegistry{
					ObjectMeta: metav1.ObjectMeta{Name: "rr1"},
					Spec: searchv1alpha1.ResourceRegistrySpec{
						TargetCluster:     policyv1alpha1.ClusterAffinity{ClusterNames: []string{"cluster1"}},
						ResourceSelectors: []searchv1alpha1.ResourceSelector{proxytest.PodSelector, proxytest.NodeSelector},
					},
				},
				&searchv1alpha1.ResourceRegistry{
					ObjectMeta: metav1.ObjectMeta{Name: "rr2"},
					Spec: searchv1alpha1.ResourceRegistrySpec{
						TargetCluster:     policyv1alpha1.ClusterAffinity{ClusterNames: []string{"cluster2"}},
						ResourceSelectors: []searchv1alpha1.ResourceSelector{proxytest.NodeSelector},
					},
				},
			},
			want: map[string]map[string]*store.MultiNamespace{
				"cluster1": {
					"pods":  newMultiNs(),
					"nodes": newMultiNs(),
				},
				"cluster2": {
					"nodes": newMultiNs(),
				},
			},
		},
		{
			name: "register pod twice in one ResourceRegistry",
			input: []runtime.Object{
				newCluster("cluster1"),
				newCluster("cluster2"),
				&searchv1alpha1.ResourceRegistry{
					ObjectMeta: metav1.ObjectMeta{Name: "rr1"},
					Spec: searchv1alpha1.ResourceRegistrySpec{
						TargetCluster: policyv1alpha1.ClusterAffinity{
							ClusterNames: []string{"cluster1"},
						},
						ResourceSelectors: []searchv1alpha1.ResourceSelector{
							proxytest.PodSelector,
							proxytest.PodSelector,
						},
					},
				},
			},
			want: map[string]map[string]*store.MultiNamespace{
				"cluster1": {
					"pods": newMultiNs(),
				},
			},
		},
		{
			name: "register pod twice in two ResourceRegistries",
			input: []runtime.Object{
				newCluster("cluster1"),
				newCluster("cluster2"),
				&searchv1alpha1.ResourceRegistry{
					ObjectMeta: metav1.ObjectMeta{Name: "rr1"},
					Spec: searchv1alpha1.ResourceRegistrySpec{
						TargetCluster: policyv1alpha1.ClusterAffinity{
							ClusterNames: []string{"cluster1"},
						},
						ResourceSelectors: []searchv1alpha1.ResourceSelector{
							proxytest.PodSelector,
						},
					},
				},
				&searchv1alpha1.ResourceRegistry{
					ObjectMeta: metav1.ObjectMeta{Name: "rr2"},
					Spec: searchv1alpha1.ResourceRegistrySpec{
						TargetCluster: policyv1alpha1.ClusterAffinity{
							ClusterNames: []string{"cluster1"},
						},
						ResourceSelectors: []searchv1alpha1.ResourceSelector{
							proxytest.PodSelector,
						},
					},
				},
			},
			want: map[string]map[string]*store.MultiNamespace{
				"cluster1": {
					"pods": newMultiNs(),
				},
			},
		},
		{
			name: "GetGroupVersionResource error shall be ignored",
			input: []runtime.Object{
				newCluster("cluster1"),
				&searchv1alpha1.ResourceRegistry{
					ObjectMeta: metav1.ObjectMeta{Name: "rr1"},
					Spec: searchv1alpha1.ResourceRegistrySpec{
						TargetCluster: policyv1alpha1.ClusterAffinity{
							ClusterNames: []string{"cluster1"},
						},
						ResourceSelectors: []searchv1alpha1.ResourceSelector{
							{APIVersion: "test.nonexist.group", Kind: "nonexist"},
						},
					},
				},
			},
			want: map[string]map[string]*store.MultiNamespace{},
		},
		{
			name: "register pod twice in two ResourceRegistries with different namespace",
			input: []runtime.Object{
				newCluster("cluster1"),
				newCluster("cluster2"),
				&searchv1alpha1.ResourceRegistry{
					ObjectMeta: metav1.ObjectMeta{Name: "rr1"},
					Spec: searchv1alpha1.ResourceRegistrySpec{
						TargetCluster: policyv1alpha1.ClusterAffinity{
							ClusterNames: []string{"cluster1"},
						},
						ResourceSelectors: []searchv1alpha1.ResourceSelector{
							proxytest.PodSelectorWithNS1,
						},
					},
				},
				&searchv1alpha1.ResourceRegistry{
					ObjectMeta: metav1.ObjectMeta{Name: "rr2"},
					Spec: searchv1alpha1.ResourceRegistrySpec{
						TargetCluster: policyv1alpha1.ClusterAffinity{
							ClusterNames: []string{"cluster1"},
						},
						ResourceSelectors: []searchv1alpha1.ResourceSelector{
							proxytest.PodSelectorWithNS2,
						},
					},
				},
			},
			want: map[string]map[string]*store.MultiNamespace{
				"cluster1": {
					"pods": newMultiNs("ns1", "ns2"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := map[string]map[string]*store.MultiNamespace{}
			karmadaClientset := karmadafake.NewSimpleClientset(tt.input...)
			karmadaFactory := karmadainformers.NewSharedInformerFactory(karmadaClientset, 0)

			ctl := &Controller{
				restMapper:     proxytest.RestMapper,
				clusterLister:  karmadaFactory.Cluster().V1alpha1().Clusters().Lister(),
				registryLister: karmadaFactory.Search().V1alpha1().ResourceRegistries().Lister(),
				store: &proxytest.MockStore{
					UpdateCacheFunc: func(m map[string]map[schema.GroupVersionResource]*store.MultiNamespace, _ map[schema.GroupVersionResource]struct{}) error {
						for clusterName, resources := range m {
							resourceCaches := map[string]*store.MultiNamespace{}
							for resource, multiNs := range resources {
								resourceCaches[resource.Resource] = multiNs
							}
							actual[clusterName] = resourceCaches
						}
						if len(actual) != len(m) {
							return fmt.Errorf("cluster duplicate: %#v", m)
						}
						return nil
					},
				},
			}
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			karmadaFactory.Start(ctx.Done())
			karmadaFactory.WaitForCacheSync(ctx.Done())

			err := ctl.reconcile(workKey)
			if err != nil {
				t.Error(err)
				return
			}
			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("diff: %v", cmp.Diff(actual, tt.want))
			}
		})
	}
}

type mockPlugin struct {
	TheOrder         int
	IsSupportRequest bool
	Called           bool
}

var _ framework.Plugin = (*mockPlugin)(nil)

func (r *mockPlugin) Order() int {
	return r.TheOrder
}

func (r *mockPlugin) SupportRequest(_ framework.ProxyRequest) bool {
	return r.IsSupportRequest
}

func (r *mockPlugin) Connect(_ context.Context, _ framework.ProxyRequest) (http.Handler, error) {
	return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		r.Called = true
	}), nil
}

func convertPluginSlice(in []*mockPlugin) []framework.Plugin {
	out := make([]framework.Plugin, 0, len(in))
	for _, plugin := range in {
		out = append(out, plugin)
	}

	return out
}

func TestController_Connect(t *testing.T) {
	store := &proxytest.MockStore{
		HasResourceFunc: func(gvr schema.GroupVersionResource) bool { return gvr == proxytest.PodGVR },
	}

	tests := []struct {
		name       string
		plugins    []*mockPlugin
		wantErr    bool
		wantCalled []bool
	}{
		{
			name: "call first",
			plugins: []*mockPlugin{
				{
					TheOrder:         0,
					IsSupportRequest: true,
				},
				{
					TheOrder:         1,
					IsSupportRequest: true,
				},
			},
			wantErr:    false,
			wantCalled: []bool{true, false},
		},
		{
			name: "call second",
			plugins: []*mockPlugin{
				{
					TheOrder:         0,
					IsSupportRequest: false,
				},
				{
					TheOrder:         1,
					IsSupportRequest: true,
				},
			},
			wantErr:    false,
			wantCalled: []bool{false, true},
		},
		{
			name: "call fail",
			plugins: []*mockPlugin{
				{
					TheOrder:         0,
					IsSupportRequest: false,
				},
				{
					TheOrder:         1,
					IsSupportRequest: false,
				},
			},
			wantErr:    true,
			wantCalled: []bool{false, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := &Controller{
				proxy:                pluginruntime.NewFramework(convertPluginSlice(tt.plugins)),
				negotiatedSerializer: scheme.Codecs.WithoutConversion(),
				store:                store,
			}

			conn, err := ctl.Connect(context.TODO(), "/api/v1/pods", nil)
			if err != nil {
				t.Fatal(err)
			}

			req, err := http.NewRequest(http.MethodGet, "/prefix/api/v1/pods", nil)
			if err != nil {
				t.Fatal(err)
			}

			recorder := httptest.NewRecorder()
			conn.ServeHTTP(recorder, req)

			response := recorder.Result()

			if (response.StatusCode != 200) != tt.wantErr {
				t.Errorf("http request returned status code = %v, want error = %v",
					response.StatusCode, tt.wantErr)
			}

			if len(tt.plugins) != len(tt.wantCalled) {
				panic("len(tt.plugins) != len(tt.wantCalled), please fix test cases")
			}

			for i, n := 0, len(tt.plugins); i < n; i++ {
				if tt.plugins[i].Called != tt.wantCalled[i] {
					t.Errorf("plugin[%v].Called = %v, want = %v", i, tt.plugins[i].Called, tt.wantCalled[i])
				}
			}
		})
	}
}

type failPlugin struct{}

var _ framework.Plugin = (*failPlugin)(nil)

func (r *failPlugin) Order() int {
	return 0
}

func (r *failPlugin) SupportRequest(_ framework.ProxyRequest) bool {
	return true
}

func (r *failPlugin) Connect(_ context.Context, _ framework.ProxyRequest) (http.Handler, error) {
	return nil, fmt.Errorf("test")
}

func TestController_Connect_Error(t *testing.T) {
	store := &proxytest.MockStore{
		HasResourceFunc: func(gvr schema.GroupVersionResource) bool {
			return gvr == proxytest.PodGVR
		},
	}

	plugins := []framework.Plugin{&failPlugin{}}

	ctl := &Controller{
		proxy:                pluginruntime.NewFramework(plugins),
		store:                store,
		negotiatedSerializer: scheme.Codecs.WithoutConversion(),
	}

	h, err := ctl.Connect(context.TODO(), "/api", nil)
	if err != nil {
		t.Error(err)
		return
	}

	response := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api", nil)
	if err != nil {
		t.Error(err)
		return
	}
	req.Header = make(http.Header)
	req.Header.Add("Accept", "application/json")
	h.ServeHTTP(response, req)
	wantBody := `{"kind":"Status","apiVersion":"get","metadata":{},"status":"Failure","message":"test","code":500}` + "\n"
	gotBody := response.Body.String()
	if wantBody != gotBody {
		t.Errorf("got body: %v", diff.StringDiff(gotBody, wantBody))
	}
}

func newCluster(name string) *clusterv1alpha1.Cluster {
	c := &clusterv1alpha1.Cluster{}
	clusterEnablements := []clusterv1alpha1.APIEnablement{
		{
			GroupVersion: "v1",
			Resources: []clusterv1alpha1.APIResource{
				{
					Kind: "Pod",
				},
			},
		},
		{
			GroupVersion: "v1",
			Resources: []clusterv1alpha1.APIResource{
				{
					Kind: "Node",
				},
			},
		},
	}
	c.Name = name
	conditions := make([]metav1.Condition, 0, 1)
	conditions = append(conditions, util.NewCondition(clusterv1alpha1.ClusterConditionReady, "", "", metav1.ConditionTrue))
	c.Status.Conditions = conditions
	c.Status.APIEnablements = clusterEnablements
	return c
}
