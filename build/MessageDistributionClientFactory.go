package build

import (
	clients1 "github.com/pip-services-users2/client-msgdistribution-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type MessageDistributionClientFactory struct {
	*cbuild.Factory
}

func NewMessageDistributionClientFactory() *MessageDistributionClientFactory {
	c := &MessageDistributionClientFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientV1Descriptor := cref.NewDescriptor("service-msgdistribution", "client", "null", "default", "1.0")
	commandableHttpClientV1Descriptor := cref.NewDescriptor("service-msgdistribution", "client", "commandable-http", "default", "1.0")
	httpClientV1Descriptor := cref.NewDescriptor("service-msgdistribution", "client", "http", "default", "1.0")
	commandableGrpcClientV1Descriptor := cref.NewDescriptor("service-msgdistribution", "client", "commandable-grpc", "default", "1.0")
	grpcClientV1Descriptor := cref.NewDescriptor("service-msgdistribution", "client", "grpc", "default", "1.0")

	c.RegisterType(nullClientV1Descriptor, clients1.NewMessageDistributionNullClientV1)
	c.RegisterType(httpClientV1Descriptor, clients1.NewMessageDistributionGrpcClientV1)
	c.RegisterType(commandableHttpClientV1Descriptor, clients1.NewMessageDistributionHttpClientV1)
	c.RegisterType(commandableGrpcClientV1Descriptor, clients1.NewMessageDistributionCommandableGrpcClientV1)
	c.RegisterType(grpcClientV1Descriptor, clients1.NewMessageDistributionGrpcClientV1)

	return c
}
