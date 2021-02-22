package user_service

import (
	"github.com/megamixparty/gobase/lib/repository"
)

// UserService struct store userDB
type UserService struct {
	db repository.IDatabaseRepository
}

// NewUserService returns service implementation of IUserService
func NewUserService(db repository.IDatabaseRepository) *UserService {
	return &UserService{db}
}
