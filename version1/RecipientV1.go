package version1

type RecipientV1 struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Language string `json:"language,omitempty"`
}

func NewEmptyRecipientV1() *RecipientV1 {
	return &RecipientV1{}
}

func NewRecipientV1(id, name, email string) *RecipientV1 {
	return &RecipientV1{
		Id:       id,
		Name:     name,
		Email:    email,
		Phone:    "",
		Language: "",
	}
}
