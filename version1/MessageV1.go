package version1

type MessageV1 struct {
	Template string `json:"template,omitempty"`
	From     string `json:"from,omitempty"`
	Cc       string `json:"cc,omitempty"`
	Bcc      string `json:"bcc,omitempty"`
	ReplyTo  string `json:"reply_to,omitempty"`
	Subject  any    `json:"subject,omitempty"`
	Text     any    `json:"text,omitempty"`
	Html     any    `json:"html,omitempty"`
}

func NewEmptyMessageV1() *MessageV1 {
	return &MessageV1{}
}
