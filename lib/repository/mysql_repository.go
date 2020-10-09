package repository

import (
	"context"

	"github.com/megamixparty/gobase/lib/model"
)

// UserRepository interface
type UserRepository interface {
	IndexUsers(context.Context) []*model.User
	GetUser(context.Context) *model.User
}
