package request

type InputUserSignIn struct {
	Email    string `json:"email" binding:"required" example:"test@test.com"`
	Password string `json:"password" binding:"required" example:"UwU"`
}
