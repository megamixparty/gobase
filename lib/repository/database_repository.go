package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/megamixparty/gobase/lib/model"
)

// IDatabaseRepository interface
type IDatabaseRepository interface {
	GetConnection(context.Context, bool) (sqlx.ExtContext, error)
	NewUserRepository(sqlx.ExtContext) IUserRepository
	Commit(sqlx.ExtContext) error
	Rollback(sqlx.ExtContext) error
}

// IUserRepository interface
type IUserRepository interface {
	SelectUsers(context.Context) ([]*model.User, error)
	SelectUser(context.Context, int64) (*model.User, error)
	CreateUser(context.Context, *model.User) error
}
