// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/karmada-io/karmada/pkg/apis/autoscaling/v1alpha1"
	scheme "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FederatedHPAsGetter has a method to return a FederatedHPAInterface.
// A group's client should implement this interface.
type FederatedHPAsGetter interface {
	FederatedHPAs(namespace string) FederatedHPAInterface
}

// FederatedHPAInterface has methods to work with FederatedHPA resources.
type FederatedHPAInterface interface {
	Create(ctx context.Context, federatedHPA *v1alpha1.FederatedHPA, opts v1.CreateOptions) (*v1alpha1.FederatedHPA, error)
	Update(ctx context.Context, federatedHPA *v1alpha1.FederatedHPA, opts v1.UpdateOptions) (*v1alpha1.FederatedHPA, error)
	UpdateStatus(ctx context.Context, federatedHPA *v1alpha1.FederatedHPA, opts v1.UpdateOptions) (*v1alpha1.FederatedHPA, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.FederatedHPA, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.FederatedHPAList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FederatedHPA, err error)
	FederatedHPAExpansion
}

// federatedHPAs implements FederatedHPAInterface
type federatedHPAs struct {
	client rest.Interface
	ns     string
}

// newFederatedHPAs returns a FederatedHPAs
func newFederatedHPAs(c *AutoscalingV1alpha1Client, namespace string) *federatedHPAs {
	return &federatedHPAs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedHPA, and returns the corresponding federatedHPA object, and an error if there is any.
func (c *federatedHPAs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FederatedHPA, err error) {
	result = &v1alpha1.FederatedHPA{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedhpas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedHPAs that match those selectors.
func (c *federatedHPAs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FederatedHPAList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FederatedHPAList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedhpas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedHPAs.
func (c *federatedHPAs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedhpas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a federatedHPA and creates it.  Returns the server's representation of the federatedHPA, and an error, if there is any.
func (c *federatedHPAs) Create(ctx context.Context, federatedHPA *v1alpha1.FederatedHPA, opts v1.CreateOptions) (result *v1alpha1.FederatedHPA, err error) {
	result = &v1alpha1.FederatedHPA{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedhpas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedHPA).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a federatedHPA and updates it. Returns the server's representation of the federatedHPA, and an error, if there is any.
func (c *federatedHPAs) Update(ctx context.Context, federatedHPA *v1alpha1.FederatedHPA, opts v1.UpdateOptions) (result *v1alpha1.FederatedHPA, err error) {
	result = &v1alpha1.FederatedHPA{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedhpas").
		Name(federatedHPA.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedHPA).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *federatedHPAs) UpdateStatus(ctx context.Context, federatedHPA *v1alpha1.FederatedHPA, opts v1.UpdateOptions) (result *v1alpha1.FederatedHPA, err error) {
	result = &v1alpha1.FederatedHPA{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedhpas").
		Name(federatedHPA.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedHPA).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the federatedHPA and deletes it. Returns an error if one occurs.
func (c *federatedHPAs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedhpas").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedHPAs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedhpas").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched federatedHPA.
func (c *federatedHPAs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FederatedHPA, err error) {
	result = &v1alpha1.FederatedHPA{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedhpas").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
