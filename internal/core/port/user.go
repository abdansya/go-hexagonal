package port

import (
	"context"

	"github.com/abdansya/go-hexagonal/internal/core/domain"
	"github.com/abdansya/go-hexagonal/internal/core/dto"
)

type UserService interface {
	Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error)
	GetByID(ctx context.Context, id uint) (*dto.UserResponse, error)
	GetByEmail(ctx context.Context, email string) (*dto.UserResponse, error)
	GetByUsername(ctx context.Context, username string) (*dto.UserResponse, error)
	List(ctx context.Context, req *dto.ListUserRequest) ([]*dto.UserResponse, int64, error)
	Update(ctx context.Context, id uint, req *dto.CreateUserRequest) (*dto.UserResponse, error)
	Delete(ctx context.Context, id uint) error
}

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByID(ctx context.Context, id uint) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByUsername(ctx context.Context, username string) (*domain.User, error)
	FindAll(ctx context.Context, req *dto.ListUserRequest) ([]*domain.User, int64, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id uint) error
}
