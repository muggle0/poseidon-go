package Model

type User struct {
	Username string `form:"username" bingding:"required"`
	Phone    string `form:"phone" bingding:"required"`
}
