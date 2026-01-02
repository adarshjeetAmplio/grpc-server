package dtos

type SignupRequest struct {
	Name string `json:"name"`;
	Email string `json:"email"`;
	Password string `json:"password"`;
}

type SignupResponse struct {
	Message string `json:"message"`
	Token string `json:"token"`;
}

type SigninRequest struct {
	Email string `json:"email"`;
	Password string `json:"password"`;
}

type SigninResponse struct {
	Mesage string `json:"message"`;
	Token string `json:"token"`;
}