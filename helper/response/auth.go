package response

type TokenResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	ID    uint64 `json:"id" `
	Email string `json:"email" `
	Name  string `json:"name" `
}
