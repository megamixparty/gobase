package user_service

import (
	"context"

	"github.com/megamixparty/gobase/lib/model"
)

// ListUsers query into repository to return users object
func (us *UserService) ListUsers(ctx context.Context) (users []*model.User, err error) {
	conn, err := us.db.GetConnection(ctx, false)
	if err != nil {
		return nil, err
	}

	userRep := us.db.NewUserRepository(conn)
	users, err = userRep.SelectUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
