package repository

import (
	"context"

	"github.com/megamixparty/gobase/lib/model"
)

// IUserRepository interface
type IUserRepository interface {
	SelectUsers(context.Context) ([]*model.User, error)
	SelectUser(context.Context, int64) (*model.User, error)
	CreateUser(context.Context, *model.User) error
}
