package fake

import (
	v1alpha1 "github.com/kyma-project/kyma/components/event-bus/generated/ea/clientset/versioned/typed/remoteenvironment.kyma.cx/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeRemoteenvironmentV1alpha1 struct {
	*testing.Fake
}

func (c *FakeRemoteenvironmentV1alpha1) EventActivations(namespace string) v1alpha1.EventActivationInterface {
	return &FakeEventActivations{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeRemoteenvironmentV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
