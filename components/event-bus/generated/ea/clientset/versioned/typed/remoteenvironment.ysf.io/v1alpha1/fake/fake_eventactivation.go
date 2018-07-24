package fake

import (
	v1alpha1 "github.com/kyma-project/kyma/components/event-bus/internal/ea/apis/remoteenvironment.kyma.cx/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEventActivations implements EventActivationInterface
type FakeEventActivations struct {
	Fake *FakeRemoteenvironmentV1alpha1
	ns   string
}

var eventactivationsResource = schema.GroupVersionResource{Group: "remoteenvironment.kyma.cx", Version: "v1alpha1", Resource: "eventactivations"}

var eventactivationsKind = schema.GroupVersionKind{Group: "remoteenvironment.kyma.cx", Version: "v1alpha1", Kind: "EventActivation"}

// Get takes name of the eventActivation, and returns the corresponding eventActivation object, and an error if there is any.
func (c *FakeEventActivations) Get(name string, options v1.GetOptions) (result *v1alpha1.EventActivation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eventactivationsResource, c.ns, name), &v1alpha1.EventActivation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventActivation), err
}

// List takes label and field selectors, and returns the list of EventActivations that match those selectors.
func (c *FakeEventActivations) List(opts v1.ListOptions) (result *v1alpha1.EventActivationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eventactivationsResource, eventactivationsKind, c.ns, opts), &v1alpha1.EventActivationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EventActivationList{}
	for _, item := range obj.(*v1alpha1.EventActivationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eventActivations.
func (c *FakeEventActivations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(eventactivationsResource, c.ns, opts))

}

// Create takes the representation of a eventActivation and creates it.  Returns the server's representation of the eventActivation, and an error, if there is any.
func (c *FakeEventActivations) Create(eventActivation *v1alpha1.EventActivation) (result *v1alpha1.EventActivation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eventactivationsResource, c.ns, eventActivation), &v1alpha1.EventActivation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventActivation), err
}

// Update takes the representation of a eventActivation and updates it. Returns the server's representation of the eventActivation, and an error, if there is any.
func (c *FakeEventActivations) Update(eventActivation *v1alpha1.EventActivation) (result *v1alpha1.EventActivation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eventactivationsResource, c.ns, eventActivation), &v1alpha1.EventActivation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventActivation), err
}

// Delete takes name of the eventActivation and deletes it. Returns an error if one occurs.
func (c *FakeEventActivations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(eventactivationsResource, c.ns, name), &v1alpha1.EventActivation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEventActivations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eventactivationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EventActivationList{})
	return err
}

// Patch applies the patch and returns the patched eventActivation.
func (c *FakeEventActivations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EventActivation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eventactivationsResource, c.ns, name, data, subresources...), &v1alpha1.EventActivation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventActivation), err
}
