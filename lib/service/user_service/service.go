package user_service

import (
	"github.com/megamixparty/gobase/lib/repository"
)

// UserService struct store userDB
type UserService struct {
	userRep repository.IUserRepository
}

// NewUserService returns service implementation of IUserService
func NewUserService(userRep repository.IUserRepository) *UserService {
	return &UserService{userRep}
}
