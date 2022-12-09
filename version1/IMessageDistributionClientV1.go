package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type IMessageDistributionClientV1 interface {
	SendMessage(ctx context.Context, correlationId string, recipient *RecipientV1,
		message *MessageV1, parameters *cconf.ConfigParams, method string) error

	SendMessages(ctx context.Context, correlationId string, recipients []*RecipientV1,
		message *MessageV1, parameters *cconf.ConfigParams, method string) error

	SendMessageToRecipient(ctx context.Context, correlationId string, recipientId string, subscription string,
		message *MessageV1, parameters *cconf.ConfigParams, method string) error

	SendMessageToRecipients(ctx context.Context, correlationId string, recipientIds []string, subscription string,
		message *MessageV1, parameters *cconf.ConfigParams, method string) error
}
