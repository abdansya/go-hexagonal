package dto

type ListUserRequest struct {
	Search  string `query:"search" validate:"omitempty,max=100"`
	Page    int    `query:"page" validate:"omitempty,min=1"`
	PerPage int    `query:"per_page" validate:"omitempty,min=5,max=200"`
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=3,alphanum"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
