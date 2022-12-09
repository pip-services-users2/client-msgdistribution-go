package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-msgdistribution-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type messageDistributionCommandaleHttpClientV1 struct {
	client  *version1.MessageDistributionHttpClientV1
	fixture *MessageDistributionClientFixtureV1
}

func newMessageDistributionCommandaleHttpClientV1() *messageDistributionCommandaleHttpClientV1 {
	return &messageDistributionCommandaleHttpClientV1{}
}

func (c *messageDistributionCommandaleHttpClientV1) setup(t *testing.T) *MessageDistributionClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewMessageDistributionHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewAccountsClientFixtureV1(c.client)

	return c.fixture
}

func (c *messageDistributionCommandaleHttpClientV1) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestHttpMessages(t *testing.T) {
	c := newMessageDistributionCommandaleHttpClientV1()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestSendMessage(t)
	fixture.TestSendMessageToRecipient(t)
}
