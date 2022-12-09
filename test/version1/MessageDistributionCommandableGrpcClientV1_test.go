package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-msgdistribution-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type messageDistributionCommandableGrpcClientV1Test struct {
	client  *version1.MessageDistributionCommandableGrpcClientV1
	fixture *MessageDistributionClientFixtureV1
}

func newMessageDistributionCommandableGrpcClientV1Test() *messageDistributionCommandableGrpcClientV1Test {
	return &messageDistributionCommandableGrpcClientV1Test{}
}

func (c *messageDistributionCommandableGrpcClientV1Test) setup(t *testing.T) *MessageDistributionClientFixtureV1 {
	var GRPC_HOST = os.Getenv("GRPC_HOST")
	if GRPC_HOST == "" {
		GRPC_HOST = "localhost"
	}
	var GRPC_PORT = os.Getenv("GRPC_PORT")
	if GRPC_PORT == "" {
		GRPC_PORT = "8090"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", GRPC_HOST,
		"connection.port", GRPC_PORT,
	)

	c.client = version1.NewMessageDistributionCommandableGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewAccountsClientFixtureV1(c.client)

	return c.fixture
}

func (c *messageDistributionCommandableGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableGrpcMessages(t *testing.T) {
	c := newMessageDistributionCommandableGrpcClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestSendMessage(t)
	fixture.TestSendMessageToRecipient(t)
}
