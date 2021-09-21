package req

type ReqSignIn struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
