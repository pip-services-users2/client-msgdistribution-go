package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cclnt "github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type MessageDistributionHttpClientV1 struct {
	*cclnt.CommandableHttpClient
	defaultParameters *cconf.ConfigParams
}

func NewMessageDistributionHttpClientV1() *MessageDistributionHttpClientV1 {
	return NewMessageDistributionHttpClientV1WithConfig(nil)
}

func NewMessageDistributionHttpClientV1WithConfig(config *cconf.ConfigParams) *MessageDistributionHttpClientV1 {
	c := &MessageDistributionHttpClientV1{}

	c.CommandableHttpClient = cclnt.NewCommandableHttpClient("v1/msg_distribution")
	if config != nil {
		c.defaultParameters = cconf.NewConfigParams(config.Value()).GetSection("parameters")
		c.Configure(context.Background(), config)
	} else {
		c.defaultParameters = cconf.NewEmptyConfigParams()
	}

	return c
}

func (c *MessageDistributionHttpClientV1) SendMessage(ctx context.Context, correlationId string, recipient *RecipientV1, message *MessageV1, parameters *cconf.ConfigParams, method string) (err error) {
	if parameters == nil {
		parameters = c.defaultParameters
	} else {
		parameters = c.defaultParameters.Override(parameters)
	}

	timing := c.Instrument(ctx, correlationId, "msg_distribution.send_message")
	defer timing.EndTiming(ctx, err)

	_, err = c.CallCommand(ctx, "send_message", correlationId, cdata.NewAnyValueMapFromTuples(
		"recipient", recipient,
		"message", message,
		"parameters", parameters,
		"method", method,
	))

	return err
}

func (c *MessageDistributionHttpClientV1) SendMessages(ctx context.Context, correlationId string, recipients []*RecipientV1, message *MessageV1, parameters *cconf.ConfigParams, method string) (err error) {
	if parameters == nil {
		parameters = c.defaultParameters
	} else {
		parameters = c.defaultParameters.Override(parameters)
	}

	timing := c.Instrument(ctx, correlationId, "msg_distribution.send_messages")
	defer timing.EndTiming(ctx, err)

	_, err = c.CallCommand(ctx, "send_messages", correlationId, cdata.NewAnyValueMapFromTuples(
		"recipient", recipients,
		"message", message,
		"parameters", parameters,
		"method", method,
	))

	return err
}

func (c *MessageDistributionHttpClientV1) SendMessageToRecipient(ctx context.Context, correlationId string, recipientId string, subscription string, message *MessageV1, parameters *cconf.ConfigParams, method string) (err error) {
	if parameters == nil {
		parameters = c.defaultParameters
	} else {
		parameters = c.defaultParameters.Override(parameters)
	}

	timing := c.Instrument(ctx, correlationId, "msg_distribution.send_message_to_recipient")
	defer timing.EndTiming(ctx, err)

	_, err = c.CallCommand(ctx, "send_message_to_recipient", correlationId, cdata.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
		"subscription", subscription,
		"message", message,
		"parameters", parameters,
		"method", method,
	))

	return err
}

func (c *MessageDistributionHttpClientV1) SendMessageToRecipients(ctx context.Context, correlationId string, recipientIds []string, subscription string, message *MessageV1, parameters *cconf.ConfigParams, method string) (err error) {
	if parameters == nil {
		parameters = c.defaultParameters
	} else {
		parameters = c.defaultParameters.Override(parameters)
	}

	timing := c.Instrument(ctx, correlationId, "msg_distribution.send_message_to_recipients")
	defer timing.EndTiming(ctx, err)

	_, err = c.CallCommand(ctx, "send_message_to_recipients", correlationId, cdata.NewAnyValueMapFromTuples(
		"recipient_ids", recipientIds,
		"subscription", subscription,
		"message", message,
		"parameters", parameters,
		"method", method,
	))

	return err
}
