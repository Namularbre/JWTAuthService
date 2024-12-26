package users

type User struct {
	IdUser      int    `json:"idUser"`
	Username    string `json:"username" form:"username" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	IdGroup     int    `json:"idGroup" form:"idGroup"`
	IsAdmin     bool   `json:"isAdmin" form:"isAdmin"`
	Email       string `json:"email" form:"email" binding:"required"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber" binding:"required"`
}
