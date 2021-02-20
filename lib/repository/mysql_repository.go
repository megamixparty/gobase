package repository

import (
	"context"

	"github.com/megamixparty/gobase/lib/model"
)

// UserRepository interface
type UserRepository interface {
	IndexUsers(context.Context) ([]*model.User, error)
	GetUser(context.Context, int64) (*model.User, error)
	CreateUser(context.Context, *model.User) error
}
