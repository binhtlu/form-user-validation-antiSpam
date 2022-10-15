package User

type user struct {
	Email       string `json:"Email" binding:"email"`
	PhoneNumber string `json:"phonenumber" binding:"required,startswith=078,len=10"`
}
type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var UserList = []user{}
