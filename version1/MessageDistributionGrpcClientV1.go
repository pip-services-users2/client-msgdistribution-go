package version1

import (
	"context"

	"github.com/pip-services-users2/client-msgdistribution-go/protos"
	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	clients "github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type MessageDistributionGrpcClientV1 struct {
	*clients.GrpcClient
	defaultParameters *cconf.ConfigParams
}

func NewMessageDistributionGrpcClientV1() *MessageDistributionGrpcClientV1 {
	return NewMessageDistributionGrpcClientV1WithConfig(nil)
}

func NewMessageDistributionGrpcClientV1WithConfig(config *cconf.ConfigParams) *MessageDistributionGrpcClientV1 {
	c := &MessageDistributionGrpcClientV1{
		GrpcClient: clients.NewGrpcClient("msgdistribution_v1.MessageDistribution"),
	}

	if config != nil {
		c.defaultParameters = cconf.NewConfigParams(config.Value()).GetSection("parameters")
		c.Configure(context.Background(), config)
	} else {
		c.defaultParameters = cconf.NewEmptyConfigParams()
	}

	return c
}

func (c *MessageDistributionGrpcClientV1) SendMessage(ctx context.Context, correlationId string, recipient *RecipientV1, message *MessageV1, parameters *cconf.ConfigParams, method string) (err error) {
	timing := c.Instrument(ctx, correlationId, "msgdistribution_v1.send_message")
	defer timing.EndTiming(ctx, err)

	if parameters == nil {
		parameters = c.defaultParameters
	} else {
		parameters = c.defaultParameters.Override(parameters)
	}

	req := &protos.SendMessageRequest{
		CorrelationId: correlationId,
		Method:        method,
		Message:       fromMessage(message),
		Parameters:    parameters.Value(),
		Recipient:     fromRecepient(recipient),
	}

	reply := new(protos.SendEmptyReply)
	err = c.CallWithContext(ctx, "send_message", correlationId, req, reply)

	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}

func (c *MessageDistributionGrpcClientV1) SendMessages(ctx context.Context, correlationId string, recipients []*RecipientV1, message *MessageV1, parameters *cconf.ConfigParams, method string) (err error) {
	timing := c.Instrument(ctx, correlationId, "msgdistribution_v1.send_messages")
	defer timing.EndTiming(ctx, err)

	if parameters == nil {
		parameters = c.defaultParameters
	} else {
		parameters = c.defaultParameters.Override(parameters)
	}

	req := &protos.SendMessagesRequest{
		CorrelationId: correlationId,
		Method:        method,
		Message:       fromMessage(message),
		Parameters:    parameters.Value(),
		Recipients:    fromRecepients(recipients),
	}

	reply := new(protos.SendEmptyReply)
	err = c.CallWithContext(ctx, "send_messages", correlationId, req, reply)

	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}

func (c *MessageDistributionGrpcClientV1) SendMessageToRecipient(ctx context.Context, correlationId string, recipientId string, subscription string, message *MessageV1, parameters *cconf.ConfigParams, method string) (err error) {
	timing := c.Instrument(ctx, correlationId, "msgdistribution_v1.send_message_to_recipient")
	defer timing.EndTiming(ctx, err)

	if parameters == nil {
		parameters = c.defaultParameters
	} else {
		parameters = c.defaultParameters.Override(parameters)
	}

	req := &protos.SendMessageWithRecipientRequest{
		CorrelationId: correlationId,
		Method:        method,
		Message:       fromMessage(message),
		Parameters:    parameters.Value(),
		Subscription:  subscription,
		RecipientId:   recipientId,
	}

	reply := new(protos.SendEmptyReply)
	err = c.CallWithContext(ctx, "send_message_to_recipient", correlationId, req, reply)

	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}

func (c *MessageDistributionGrpcClientV1) SendMessageToRecipients(ctx context.Context, correlationId string, recipientIds []string, subscription string, message *MessageV1, parameters *cconf.ConfigParams, method string) (err error) {
	timing := c.Instrument(ctx, correlationId, "msgdistribution_v1.send_message_to_recipients")
	defer timing.EndTiming(ctx, err)

	if parameters == nil {
		parameters = c.defaultParameters
	} else {
		parameters = c.defaultParameters.Override(parameters)
	}

	req := &protos.SendMessageWithRecipientsRequest{
		CorrelationId: correlationId,
		Method:        method,
		Message:       fromMessage(message),
		Parameters:    parameters.Value(),
		Subscription:  subscription,
		RecipientIds:  recipientIds,
	}

	reply := new(protos.SendEmptyReply)
	err = c.CallWithContext(ctx, "send_message_to_recipients", correlationId, req, reply)

	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}
