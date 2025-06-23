package mapper

import (
	"time"

	"github.com/abdansya/go-hexagonal/internal/core/domain"
	"github.com/abdansya/go-hexagonal/internal/core/dto"
)

func ToUserResponse(user *domain.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}

func ToUserResponseList(users []*domain.User) []*dto.UserResponse {
	var responses []*dto.UserResponse
	for _, user := range users {
		responses = append(responses, ToUserResponse(user))
	}
	return responses
}
