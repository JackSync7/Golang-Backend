package usersdto

type CreateUserRequest struct {
	ID        int    `json:"id"`
	Fullname  string `json:"fullname" form:"name" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required"`
	Password  string `json:"password" form:"password" validate:"required"`
	Gender    string `json:"gender" form:"gender" validate:"required"`
	Phone     string `json:"phone" form:"phone" validate:"required"`
	Address   string `json:"address" form:"address" validate:"required"`
	Subscribe bool   `json:"subscribe" form:"false"`
}

type UpdateUserRequest struct {
	Fullname  string `json:"fullname" form:"name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Gender    string `json:"gender" form:"gender" validate:"required"`
	Phone     string `json:"phone" form:"phone" validate:"required"`
	Address   string `json:"address" form:"address" validate:"required"`
	Subscribe bool   `json:"subscribe" form:"false"`
}
