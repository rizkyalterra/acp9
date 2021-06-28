package user

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Passowrd string `json:"password"`
	School   School `json:"school"`
}
