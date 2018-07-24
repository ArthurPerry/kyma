// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kyma-project/kyma/components/event-bus/api/push/eventing.kyma.cx/v1alpha1"
	scheme "github.com/kyma-project/kyma/components/event-bus/generated/push/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SubscriptionsGetter has a method to return a SubscriptionInterface.
// A group's client should implement this interface.
type SubscriptionsGetter interface {
	Subscriptions(namespace string) SubscriptionInterface
}

// SubscriptionInterface has methods to work with Subscription resources.
type SubscriptionInterface interface {
	Create(*v1alpha1.Subscription) (*v1alpha1.Subscription, error)
	Update(*v1alpha1.Subscription) (*v1alpha1.Subscription, error)
	UpdateStatus(*v1alpha1.Subscription) (*v1alpha1.Subscription, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Subscription, error)
	List(opts v1.ListOptions) (*v1alpha1.SubscriptionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Subscription, err error)
	SubscriptionExpansion
}

// subscriptions implements SubscriptionInterface
type subscriptions struct {
	client rest.Interface
	ns     string
}

// newSubscriptions returns a Subscriptions
func newSubscriptions(c *EventingV1alpha1Client, namespace string) *subscriptions {
	return &subscriptions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the subscription, and returns the corresponding subscription object, and an error if there is any.
func (c *subscriptions) Get(name string, options v1.GetOptions) (result *v1alpha1.Subscription, err error) {
	result = &v1alpha1.Subscription{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("subscriptions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Subscriptions that match those selectors.
func (c *subscriptions) List(opts v1.ListOptions) (result *v1alpha1.SubscriptionList, err error) {
	result = &v1alpha1.SubscriptionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("subscriptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested subscriptions.
func (c *subscriptions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("subscriptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a subscription and creates it.  Returns the server's representation of the subscription, and an error, if there is any.
func (c *subscriptions) Create(subscription *v1alpha1.Subscription) (result *v1alpha1.Subscription, err error) {
	result = &v1alpha1.Subscription{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("subscriptions").
		Body(subscription).
		Do().
		Into(result)
	return
}

// Update takes the representation of a subscription and updates it. Returns the server's representation of the subscription, and an error, if there is any.
func (c *subscriptions) Update(subscription *v1alpha1.Subscription) (result *v1alpha1.Subscription, err error) {
	result = &v1alpha1.Subscription{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("subscriptions").
		Name(subscription.Name).
		Body(subscription).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *subscriptions) UpdateStatus(subscription *v1alpha1.Subscription) (result *v1alpha1.Subscription, err error) {
	result = &v1alpha1.Subscription{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("subscriptions").
		Name(subscription.Name).
		SubResource("status").
		Body(subscription).
		Do().
		Into(result)
	return
}

// Delete takes name of the subscription and deletes it. Returns an error if one occurs.
func (c *subscriptions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("subscriptions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *subscriptions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("subscriptions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched subscription.
func (c *subscriptions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Subscription, err error) {
	result = &v1alpha1.Subscription{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("subscriptions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
