package user_service

import (
	"github.com/megamixparty/gobase/lib/repository"
)

// UserService struct store userDB
type UserService struct {
	userRep repository.UserRepository
}

// NewUserService returns service implementation of UserService
func NewUserService(userRep repository.UserRepository) *UserService {
	return &UserService{userRep}
}
