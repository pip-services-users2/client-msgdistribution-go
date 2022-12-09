package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	clients "github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type MessageDistributionCommandableGrpcClientV1 struct {
	*clients.CommandableGrpcClient
	defaultParameters *cconf.ConfigParams
}

func NewMessageDistributionCommandableGrpcClientV1() *MessageDistributionCommandableGrpcClientV1 {
	return NewMessageDistributionCommandableGrpcClientV1WithConfig(nil)
}

func NewMessageDistributionCommandableGrpcClientV1WithConfig(config *cconf.ConfigParams) *MessageDistributionCommandableGrpcClientV1 {
	c := &MessageDistributionCommandableGrpcClientV1{
		CommandableGrpcClient: clients.NewCommandableGrpcClient("v1/msg_distribution"),
	}

	if config != nil {
		c.defaultParameters = cconf.NewConfigParams(config.Value()).GetSection("parameters")
		c.Configure(context.Background(), config)
	} else {
		c.defaultParameters = cconf.NewEmptyConfigParams()
	}

	return c
}

func (c *MessageDistributionCommandableGrpcClientV1) SendMessage(ctx context.Context, correlationId string, recipient *RecipientV1, message *MessageV1, parameters *cconf.ConfigParams, method string) (err error) {
	timing := c.Instrument(ctx, correlationId, "msgdistribution_v1.send_message")
	defer timing.EndTiming(ctx, err)

	if parameters == nil {
		parameters = c.defaultParameters
	} else {
		parameters = c.defaultParameters.Override(parameters)
	}

	params := cdata.NewAnyValueMapFromTuples(
		"recipient", recipient,
		"message", message,
		"parameters", parameters,
		"method", method,
	)

	_, err = c.CallCommand(ctx, "send_message", correlationId, params)

	return err
}

func (c *MessageDistributionCommandableGrpcClientV1) SendMessages(ctx context.Context, correlationId string, recipients []*RecipientV1, message *MessageV1, parameters *cconf.ConfigParams, method string) (err error) {
	timing := c.Instrument(ctx, correlationId, "msgdistribution_v1.send_messages")
	defer timing.EndTiming(ctx, err)

	if parameters == nil {
		parameters = c.defaultParameters
	} else {
		parameters = c.defaultParameters.Override(parameters)
	}

	params := cdata.NewAnyValueMapFromTuples(
		"recipients", recipients,
		"message", message,
		"parameters", parameters,
		"method", method,
	)

	_, err = c.CallCommand(ctx, "send_messages", correlationId, params)

	return err
}

func (c *MessageDistributionCommandableGrpcClientV1) SendMessageToRecipient(ctx context.Context, correlationId string, recipientId string, subscription string, message *MessageV1, parameters *cconf.ConfigParams, method string) (err error) {
	timing := c.Instrument(ctx, correlationId, "msgdistribution_v1.send_message_to_recipient")
	defer timing.EndTiming(ctx, err)

	if parameters == nil {
		parameters = c.defaultParameters
	} else {
		parameters = c.defaultParameters.Override(parameters)
	}

	params := cdata.NewAnyValueMapFromTuples(
		"recipient_id", recipientId,
		"subscription", subscription,
		"message", message,
		"parameters", parameters,
		"method", method,
	)

	_, err = c.CallCommand(ctx, "send_message_to_recipient", correlationId, params)

	return err
}

func (c *MessageDistributionCommandableGrpcClientV1) SendMessageToRecipients(ctx context.Context, correlationId string, recipientIds []string, subscription string, message *MessageV1, parameters *cconf.ConfigParams, method string) (err error) {
	timing := c.Instrument(ctx, correlationId, "msgdistribution_v1.send_message_to_recipients")
	defer timing.EndTiming(ctx, err)

	if parameters == nil {
		parameters = c.defaultParameters
	} else {
		parameters = c.defaultParameters.Override(parameters)
	}

	params := cdata.NewAnyValueMapFromTuples(
		"recipient_ids", recipientIds,
		"subscription", subscription,
		"message", message,
		"parameters", parameters,
		"method", method,
	)

	_, err = c.CallCommand(ctx, "send_message_to_recipients", correlationId, params)

	return err
}
