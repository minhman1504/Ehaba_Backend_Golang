package req

type ReqSignUp struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`

	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"` // tags
	Gender    string `json:"gender,omitempty"`

	Birthday string `json:"birthday,omitempty"`
}
