package models



type LoginRequest struct {
	User UserLogin `json:"user"`
}

type UserLogin struct {
	Email string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

/*
func ValidateUser(email, password string) bool {

	if email == "admin@picklejar.com" && password == "123" ||
		email == "user1" && password == "user1" ||
		email == "user2" && password == "user2" {
		return true
	}
	return false


	return true
}
*/
