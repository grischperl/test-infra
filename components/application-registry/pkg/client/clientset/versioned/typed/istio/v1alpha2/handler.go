// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	"time"

	v1alpha2 "github.com/kyma-project/kyma/components/application-registry/pkg/apis/istio/v1alpha2"
	scheme "github.com/kyma-project/kyma/components/application-registry/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HandlersGetter has a method to return a HandlerInterface.
// A group's client should implement this interface.
type HandlersGetter interface {
	Handlers(namespace string) HandlerInterface
}

// HandlerInterface has methods to work with Handler resources.
type HandlerInterface interface {
	Create(*v1alpha2.Handler) (*v1alpha2.Handler, error)
	Update(*v1alpha2.Handler) (*v1alpha2.Handler, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.Handler, error)
	List(opts v1.ListOptions) (*v1alpha2.HandlerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.Handler, err error)
	HandlerExpansion
}

// handlers implements HandlerInterface
type handlers struct {
	client rest.Interface
	ns     string
}

// newHandlers returns a Handlers
func newHandlers(c *IstioV1alpha2Client, namespace string) *handlers {
	return &handlers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the handler, and returns the corresponding handler object, and an error if there is any.
func (c *handlers) Get(name string, options v1.GetOptions) (result *v1alpha2.Handler, err error) {
	result = &v1alpha2.Handler{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("handlers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Handlers that match those selectors.
func (c *handlers) List(opts v1.ListOptions) (result *v1alpha2.HandlerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.HandlerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("handlers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested handlers.
func (c *handlers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("handlers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a handler and creates it.  Returns the server's representation of the handler, and an error, if there is any.
func (c *handlers) Create(handler *v1alpha2.Handler) (result *v1alpha2.Handler, err error) {
	result = &v1alpha2.Handler{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("handlers").
		Body(handler).
		Do().
		Into(result)
	return
}

// Update takes the representation of a handler and updates it. Returns the server's representation of the handler, and an error, if there is any.
func (c *handlers) Update(handler *v1alpha2.Handler) (result *v1alpha2.Handler, err error) {
	result = &v1alpha2.Handler{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("handlers").
		Name(handler.Name).
		Body(handler).
		Do().
		Into(result)
	return
}

// Delete takes name of the handler and deletes it. Returns an error if one occurs.
func (c *handlers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("handlers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *handlers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("handlers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched handler.
func (c *handlers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.Handler, err error) {
	result = &v1alpha2.Handler{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("handlers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}