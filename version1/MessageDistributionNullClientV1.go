package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type MessageDistributionNullClientV1 struct {
}

func NewMessageDistributionNullClientV1() *MessageDistributionNullClientV1 {
	return &MessageDistributionNullClientV1{}
}

func (c *MessageDistributionNullClientV1) SendMessage(ctx context.Context, correlationId string, recipient *RecipientV1, message *MessageV1, parameters *cconf.ConfigParams, method string) error {
	return nil
}

func (c *MessageDistributionNullClientV1) SendMessages(ctx context.Context, correlationId string, recipients []*RecipientV1, message *MessageV1, parameters *cconf.ConfigParams, method string) error {
	return nil
}

func (c *MessageDistributionNullClientV1) SendMessageToRecipient(ctx context.Context, correlationId string, recipientId string, subscription string, message *MessageV1, parameters *cconf.ConfigParams, method string) error {
	return nil
}

func (c *MessageDistributionNullClientV1) SendMessageToRecipients(ctx context.Context, correlationId string, recipientIds []string, subscription string, message *MessageV1, parameters *cconf.ConfigParams, method string) error {
	return nil
}
