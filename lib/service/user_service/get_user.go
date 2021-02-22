package user_service

import (
	"context"

	"github.com/megamixparty/gobase/lib/model"
)

// GetUser query into repository to return users object
func (us *UserService) GetUser(ctx context.Context, id int64) (user *model.User, err error) {
	conn, err := us.db.GetConnection(ctx, false)
	if err != nil {
		return nil, err
	}

	userRep := us.db.NewUserRepository(conn)
	user, err = userRep.SelectUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
