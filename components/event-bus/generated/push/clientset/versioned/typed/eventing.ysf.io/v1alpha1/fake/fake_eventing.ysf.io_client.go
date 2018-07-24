package fake

import (
	v1alpha1 "github.com/kyma-project/kyma/components/event-bus/generated/push/clientset/versioned/typed/eventing.kyma.cx/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeEventingV1alpha1 struct {
	*testing.Fake
}

func (c *FakeEventingV1alpha1) Subscriptions(namespace string) v1alpha1.SubscriptionInterface {
	return &FakeSubscriptions{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeEventingV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
