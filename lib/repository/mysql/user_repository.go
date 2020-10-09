package mysql

import (
	"context"
	"database/sql"

	"github.com/megamixparty/gobase/lib/model"
)

const (
	tableName = "users"
)

// UserMysql struct contains Database to access User table
type UserMysql struct {
	db *sql.DB
}

// NewUserMysql returns database implementation of UserRepository
func NewUserMysql(db *sql.DB) *UserMysql {
	return &UserMysql{db}
}

// GetUser get user object based on id
func (um *UserMysql) GetUser(ctx context.Context, id int64) (user *model.User, err error) {
	err = um.db.QueryRow("SELECT id, first_name, last_name FROM ? WHERE id = ?", tableName, id).Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}

	return user, err
}
