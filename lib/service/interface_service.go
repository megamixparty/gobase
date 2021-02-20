package service

import (
	"context"

	"github.com/megamixparty/gobase/lib/model"
)

// IUserService interface
type IUserService interface {
	ListUsers(context.Context) ([]*model.User, error)
	GetUser(context.Context, int64) (*model.User, error)
	CreateUser(context.Context, *model.User) error
}
