package web

type UserCreateRequest struct {
	FirstName string `validate:"required,min=1,max=25" json:"first_name"`
	LastName  string `validate:"required,min=1,max=25" json:"last_name"`
	Email     string `validate:"required,email" json:"email"`
	Password  string `validate:"required" json:"password"`
}
