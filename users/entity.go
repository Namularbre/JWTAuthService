package users

type User struct {
	IdUser   int    `json:"idUser"`
	Username string `json:"username"`
	Password []byte `json:"password"`
}
