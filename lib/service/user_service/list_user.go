package user_service

import (
	"context"

	"github.com/megamixparty/gobase/lib/model"
)

// ListUser query into repository to return users object
func (us *UserService) ListUser(ctx context.Context) (users []*model.User, err error) {
	users, err = us.userRep.SelectUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
