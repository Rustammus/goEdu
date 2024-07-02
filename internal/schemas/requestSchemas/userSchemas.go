package requestSchemas

type InputUserSignIn struct {
	Email    string `json:"email" binding:"required" example:"test@test.com" format:"email"`
	Password string `json:"password" binding:"required" example:"UwU"`
}
