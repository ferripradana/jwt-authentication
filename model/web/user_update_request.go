package web

type UserUpdateRequest struct {
	Id        string `validate:"required" json:"id"`
	FirstName string `validate:"required,min=1,max=15" json:"first_name"`
	LastName  string `validate:"required,min=1,max=15" json:"last_name"`
}
