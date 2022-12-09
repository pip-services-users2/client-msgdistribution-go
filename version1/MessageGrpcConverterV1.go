package version1

import (
	"encoding/json"

	"github.com/pip-services-users2/client-msgdistribution-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
)

func fromError(err error) *protos.ErrorDescription {
	if err == nil {
		return nil
	}

	desc := errors.ErrorDescriptionFactory.Create(err)
	obj := &protos.ErrorDescription{
		Type:          desc.Type,
		Category:      desc.Category,
		Code:          desc.Code,
		CorrelationId: desc.CorrelationId,
		Status:        convert.StringConverter.ToString(desc.Status),
		Message:       desc.Message,
		Cause:         desc.Cause,
		StackTrace:    desc.StackTrace,
		Details:       fromMap(desc.Details),
	}

	return obj
}

func toError(obj *protos.ErrorDescription) error {
	if obj == nil || (obj.Category == "" && obj.Message == "") {
		return nil
	}

	description := &errors.ErrorDescription{
		Type:          obj.Type,
		Category:      obj.Category,
		Code:          obj.Code,
		CorrelationId: obj.CorrelationId,
		Status:        convert.IntegerConverter.ToInteger(obj.Status),
		Message:       obj.Message,
		Cause:         obj.Cause,
		StackTrace:    obj.StackTrace,
		Details:       toMap(obj.Details),
	}

	return errors.ApplicationErrorFactory.Create(description)
}

func fromMap(val map[string]interface{}) map[string]string {
	r := map[string]string{}

	for k, v := range val {
		r[k] = convert.StringConverter.ToString(v)
	}

	return r
}

func toMap(val map[string]string) map[string]interface{} {
	var r map[string]interface{}

	for k, v := range val {
		r[k] = v
	}

	return r
}

func toJson(value interface{}) string {
	if value == nil {
		return ""
	}

	b, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(b[:])
}

func fromJson(value string) interface{} {
	if value == "" {
		return nil
	}

	var m interface{}
	json.Unmarshal([]byte(value), &m)
	return m
}

func fromMessage(message *MessageV1) *protos.Message {
	if message == nil {
		return nil
	}

	obj := &protos.Message{
		Template: message.Template,
		From:     message.From,
		Cc:       message.Cc,
		Bcc:      message.Bcc,
		ReplyTo:  message.ReplyTo,
		Subject:  toJson(message.Subject),
		Text:     toJson(message.Text),
		Html:     toJson(message.Html),
	}

	return obj
}

func toMessage(obj *protos.Message) *MessageV1 {
	if obj == nil {
		return nil
	}

	msg := &MessageV1{
		Template: obj.Template,
		From:     obj.From,
		Cc:       obj.Cc,
		Bcc:      obj.Bcc,
		ReplyTo:  obj.ReplyTo,
		Subject:  obj.Subject,
		Text:     obj.Text,
		Html:     obj.Html,
	}

	return msg
}

func fromRecepients(recepients []*RecipientV1) []*protos.Recipient {
	if recepients == nil {
		return nil
	}

	obj := make([]*protos.Recipient, len(recepients))

	for _, recepient := range recepients {
		obj = append(obj, fromRecepient(recepient))
	}

	return obj
}

func toRecepients(objects []*protos.Recipient) []*RecipientV1 {
	if objects == nil {
		return nil
	}

	recepients := make([]*RecipientV1, len(objects))

	for _, obj := range objects {
		recepients = append(recepients, toRecepient(obj))
	}

	return recepients
}

func fromRecepient(recepient *RecipientV1) *protos.Recipient {
	if recepient == nil {
		return nil
	}

	obj := &protos.Recipient{
		Id:       recepient.Id,
		Name:     recepient.Name,
		Email:    recepient.Email,
		Phone:    recepient.Phone,
		Language: recepient.Language,
	}

	return obj
}

func toRecepient(obj *protos.Recipient) *RecipientV1 {
	if obj == nil {
		return nil
	}

	recepient := &RecipientV1{
		Id:       obj.Id,
		Name:     obj.Name,
		Email:    obj.Email,
		Phone:    obj.Phone,
		Language: obj.Language,
	}

	return recepient
}
