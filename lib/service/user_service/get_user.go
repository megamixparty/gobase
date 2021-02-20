package user_service

import (
	"context"

	"github.com/megamixparty/gobase/lib/model"
)

// GetUser query into repository to return users object
func (us *UserService) GetUser(ctx context.Context, id int64) (user *model.User, err error) {
	user, err = us.userRep.SelectUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
