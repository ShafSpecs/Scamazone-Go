package requests

type CustomerLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
