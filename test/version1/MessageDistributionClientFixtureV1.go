package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-users2/client-msgdistribution-go/version1"
	"github.com/stretchr/testify/assert"
)

type MessageDistributionClientFixtureV1 struct {
	Client version1.IMessageDistributionClientV1
}

func NewAccountsClientFixtureV1(client version1.IMessageDistributionClientV1) *MessageDistributionClientFixtureV1 {
	return &MessageDistributionClientFixtureV1{
		Client: client,
	}
}

func (c *MessageDistributionClientFixtureV1) TestSendMessageToRecipient(t *testing.T) {
	message := &version1.MessageV1{
		Subject: "Test subject",
		Text:    "Test text",
		Html:    "Test html",
	}

	err := c.Client.SendMessageToRecipient(context.Background(), "123", "1", "", message, nil, version1.All)

	assert.Nil(t, err)
}

func (c *MessageDistributionClientFixtureV1) TestSendMessage(t *testing.T) {
	message := &version1.MessageV1{
		Subject: "Test subject",
		Text:    "Test text",
		Html:    "Test html",
	}

	recipient := &version1.RecipientV1{
		Name:  "Test user",
		Email: "somebody@somewhere.com",
		Phone: "+12345349458",
	}

	err := c.Client.SendMessage(context.Background(), "123", recipient, message, nil, version1.All)

	assert.Nil(t, err)
}
