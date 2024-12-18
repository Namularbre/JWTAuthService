package users

type User struct {
	IdUser   int    `json:"idUser"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	IdGroup  int    `json:"idGroup" form:"idGroup"`
	IsAdmin  bool   `json:"isAdmin" form:"isAdmin"`
}
