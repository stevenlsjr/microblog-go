package api

type CreateUserParams struct {
	Username string `json:"username" form:"username" xml:"username"`
	Email    string `form:"email" json:"email" xml:"email" binding:"required"`
}
