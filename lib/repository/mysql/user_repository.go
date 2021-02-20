package mysql

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/megamixparty/gobase/lib/model"
)

const (
	userTableName = "users"
)

// UserMysql struct contains Database to access User table
type UserMysql struct {
	db sqlx.ExtContext
}

// NewUserMysql returns database implementation of UserRepository
func NewUserMysql(db sqlx.ExtContext) *UserMysql {
	return &UserMysql{db}
}

// SelectUsers do SELECT query to return user object
func (um *UserMysql) SelectUsers(ctx context.Context) (users []*model.User, err error) {
	query := fmt.Sprintf(`
		SELECT id, first_name, last_name
		FROM %s
		`, userTableName)
	rows, err := um.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := model.User{}
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, err
}

// SelectUser do SELECT query to return user object based on id
func (um *UserMysql) SelectUser(ctx context.Context, id int64) (user *model.User, err error) {
	query := fmt.Sprintf(`
		SELECT id, first_name, last_name
		FROM %s
		WHERE id = ?
		`, userTableName)
	err = um.db.QueryRowxContext(ctx, query, id).Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateUser do INSERT query to user object into database
func (um *UserMysql) CreateUser(ctx context.Context, user *model.User) (err error) {
	query := fmt.Sprintf(`
		INSERT INTO %s
		(first_name, last_name)
		VALUES
		(?,?)
		RETURNING ID
		`, userTableName)
	rows, err := um.db.QueryxContext(ctx, query, user.FirstName, user.LastName)
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID)
		if err != nil {
			return nil
		}
	}

	return nil
}
