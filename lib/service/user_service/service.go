package user_service

import (
	"github.com/megamixparty/gobase/lib/repository"
)

// UserService struct store userDB
type UserService struct {
	db    repository.IDatabaseRepository
	redis repository.IMemoryRepository
}

// NewUserService returns service implementation of IUserService
func NewUserService(db repository.IDatabaseRepository, redis repository.IMemoryRepository) *UserService {
	return &UserService{db, redis}
}
