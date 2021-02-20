package handler

import (
	"encoding/json"
	"net/http"

	"github.com/megamixparty/gobase/lib/model"
	"github.com/megamixparty/gobase/lib/service"
)

// UserHandler struct store User Service
type UserHandler struct {
	userServ service.IUserService
}

// NewUserHandler returns UserHandler
func NewUserHandler(userServ service.IUserService) *UserHandler {
	return &UserHandler{userServ}
}

// GetUsers handles '/users' endpoint
func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := uh.userServ.ListUsers(ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Maaf terjadi kesalahan",
			Meta: model.MetaResponse{
				HTTPStatus: http.StatusBadRequest,
			},
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{
		Data: users,
		Meta: model.MetaResponse{
			HTTPStatus: http.StatusBadRequest,
		},
	})

}
