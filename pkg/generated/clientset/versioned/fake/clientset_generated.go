/*
Copyright The Karmada Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/karmada-io/karmada/pkg/generated/clientset/versioned"
	autoscalingv1alpha1 "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/typed/autoscaling/v1alpha1"
	fakeautoscalingv1alpha1 "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/typed/autoscaling/v1alpha1/fake"
	clusterv1alpha1 "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/typed/cluster/v1alpha1"
	fakeclusterv1alpha1 "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/typed/cluster/v1alpha1/fake"
	configv1alpha1 "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/typed/config/v1alpha1"
	fakeconfigv1alpha1 "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/typed/config/v1alpha1/fake"
	networkingv1alpha1 "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/typed/networking/v1alpha1"
	fakenetworkingv1alpha1 "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/typed/networking/v1alpha1/fake"
	policyv1alpha1 "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/typed/policy/v1alpha1"
	fakepolicyv1alpha1 "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/typed/policy/v1alpha1/fake"
	searchv1alpha1 "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/typed/search/v1alpha1"
	fakesearchv1alpha1 "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/typed/search/v1alpha1/fake"
	workv1alpha1 "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/typed/work/v1alpha1"
	fakeworkv1alpha1 "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/typed/work/v1alpha1/fake"
	workv1alpha2 "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/typed/work/v1alpha2"
	fakeworkv1alpha2 "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/typed/work/v1alpha2/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// AutoscalingV1alpha1 retrieves the AutoscalingV1alpha1Client
func (c *Clientset) AutoscalingV1alpha1() autoscalingv1alpha1.AutoscalingV1alpha1Interface {
	return &fakeautoscalingv1alpha1.FakeAutoscalingV1alpha1{Fake: &c.Fake}
}

// ClusterV1alpha1 retrieves the ClusterV1alpha1Client
func (c *Clientset) ClusterV1alpha1() clusterv1alpha1.ClusterV1alpha1Interface {
	return &fakeclusterv1alpha1.FakeClusterV1alpha1{Fake: &c.Fake}
}

// ConfigV1alpha1 retrieves the ConfigV1alpha1Client
func (c *Clientset) ConfigV1alpha1() configv1alpha1.ConfigV1alpha1Interface {
	return &fakeconfigv1alpha1.FakeConfigV1alpha1{Fake: &c.Fake}
}

// NetworkingV1alpha1 retrieves the NetworkingV1alpha1Client
func (c *Clientset) NetworkingV1alpha1() networkingv1alpha1.NetworkingV1alpha1Interface {
	return &fakenetworkingv1alpha1.FakeNetworkingV1alpha1{Fake: &c.Fake}
}

// PolicyV1alpha1 retrieves the PolicyV1alpha1Client
func (c *Clientset) PolicyV1alpha1() policyv1alpha1.PolicyV1alpha1Interface {
	return &fakepolicyv1alpha1.FakePolicyV1alpha1{Fake: &c.Fake}
}

// SearchV1alpha1 retrieves the SearchV1alpha1Client
func (c *Clientset) SearchV1alpha1() searchv1alpha1.SearchV1alpha1Interface {
	return &fakesearchv1alpha1.FakeSearchV1alpha1{Fake: &c.Fake}
}

// WorkV1alpha1 retrieves the WorkV1alpha1Client
func (c *Clientset) WorkV1alpha1() workv1alpha1.WorkV1alpha1Interface {
	return &fakeworkv1alpha1.FakeWorkV1alpha1{Fake: &c.Fake}
}

// WorkV1alpha2 retrieves the WorkV1alpha2Client
func (c *Clientset) WorkV1alpha2() workv1alpha2.WorkV1alpha2Interface {
	return &fakeworkv1alpha2.FakeWorkV1alpha2{Fake: &c.Fake}
}
