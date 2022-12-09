package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cconv "github.com/pip-services3-gox/pip-services3-commons-gox/convert"
)

type MessageResolverV1 struct {
	config   *cconf.ConfigParams
	messages map[string]*MessageV1
}

func NewMessageResolverV1() *MessageResolverV1 {
	return NewMessageResolverV1WithConfig(nil)
}

func NewMessageResolverV1WithConfig(config *cconf.ConfigParams) *MessageResolverV1 {
	c := &MessageResolverV1{
		messages: make(map[string]*MessageV1, 0),
		config:   cconf.NewEmptyConfigParams(),
	}

	if config != nil {
		c.Configure(context.Background(), config)
	}

	return c
}

func NewMessageResolverV1FromTuples(tuples ...any) *MessageResolverV1 {
	result := NewMessageResolverV1()
	if len(tuples) == 0 {
		return result
	}

	for index := 0; index < len(tuples); index += 2 {
		if index+1 >= len(tuples) {
			break
		}

		name := cconv.StringConverter.ToString(tuples[index])
		templates := tuples[index+1]

		result.Put(name, templates)
	}

	return result

}

func (c *MessageResolverV1) Configure(ctx context.Context, config *cconf.ConfigParams) {
	c.config = config.GetSection("message_templates")
}

func (c *MessageResolverV1) Put(name string, template any) {
	c.config.Put(name, template)
}

func (c *MessageResolverV1) Resolve(name string) *MessageV1 {
	if name == "" {
		return nil
	}

	// Retrieve template first
	message, messageOk := c.messages[name]
	if messageOk {
		return message
	}

	template, templateOk := c.config.Get(name)
	if strTemp, ok := template.(string); templateOk && ok && strTemp != "" {
		// Construct a message
		message := &MessageV1{
			Template: strTemp,
		}

		// Cache the message
		c.messages[name] = message

		return message
	} else {
		// Get configuration
		config := c.config.GetSection(name)

		// Construct a message
		template := config.GetAsString("template")
		subject, subjectOk := config.GetAsObject("subject")
		text, textOk := config.GetAsObject("text")
		html, htmlOk := config.GetAsObject("html")

		message := &MessageV1{
			Template: template,
			Subject:  subject,
			Text:     text,
			Html:     html,
		}

		// Check and cache the message
		if template != "" || subjectOk || textOk || htmlOk {
			c.messages[name] = message
		} else {
			message = nil
		}

		return message
	}
}
