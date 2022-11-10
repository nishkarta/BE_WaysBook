package usersdto

type UpdateUserRequest struct {
	FullName string `json:"fullName" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Address  string `json:"address"`
	Image    string `json:"image" form:"image"`
}
