package service

import (
	"context"
	"time"

	"github.com/abdansya/go-hexagonal/internal/core/domain"
	"github.com/abdansya/go-hexagonal/internal/core/dto"
	"github.com/abdansya/go-hexagonal/internal/core/dto/mapper"
	"github.com/abdansya/go-hexagonal/internal/core/port"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) port.UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	user := &domain.User{
		Name:      req.Name,
		Email:     req.Email,
		Username:  req.Username,
		Password:  req.Password, // In production, hash the password!
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}
	return mapper.ToUserResponse(user), nil
}

func (s *UserService) GetByID(ctx context.Context, id uint) (*dto.UserResponse, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.ToUserResponse(user), nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*dto.UserResponse, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return mapper.ToUserResponse(user), nil
}

func (s *UserService) GetByUsername(ctx context.Context, username string) (*dto.UserResponse, error) {
	user, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return mapper.ToUserResponse(user), nil
}

func (s *UserService) List(ctx context.Context, req *dto.ListUserRequest) ([]*dto.UserResponse, int64, error) {
	users, total, err := s.repo.FindAll(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	return mapper.ToUserResponseList(users), total, nil
}

func (s *UserService) Update(ctx context.Context, id uint, req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	user.Name = req.Name
	user.Email = req.Email
	user.Username = req.Username
	user.Password = req.Password // In production, hash the password!
	user.UpdatedAt = time.Now()
	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}
	return mapper.ToUserResponse(user), nil
}

func (s *UserService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
