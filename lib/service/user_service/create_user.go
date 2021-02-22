package user_service

import (
	"context"
	"errors"
	"strings"

	"github.com/megamixparty/gobase/lib/model"
)

// CreateUser query into repository to return users object
func (us *UserService) CreateUser(ctx context.Context, user *model.User) (err error) {
	err = us.validateCreateUser(user)
	if err != nil {
		return err
	}

	conn, err := us.db.GetConnection(ctx, false)
	if err != nil {
		return err
	}

	userRep := us.db.NewUserRepository(conn)
	err = userRep.CreateUser(ctx, user)
	if err != nil {
		us.db.Rollback(conn)
		return err
	}

	err = us.db.Commit(conn)
	if err != nil {
		return err
	}

	return nil
}

// validate validates user object and return error if not satisfied
func (us *UserService) validateCreateUser(user *model.User) (err error) {
	errs := []string{}

	if user.FirstName == "" {
		errs = append(errs, "first_name: nama depan tidak boleh kosong")
	}
	if user.LastName == "" {
		errs = append(errs, "last_name: nama belakang tidak boleh kosong")
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ", "))
	}

	return nil
}
