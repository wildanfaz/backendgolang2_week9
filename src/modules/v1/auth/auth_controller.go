package auth

import (
	"encoding/json"
	"net/http"

	"github.com/wildanfaz/backendgolang2_week9/src/database/orm/models"
	"github.com/wildanfaz/backendgolang2_week9/src/interfaces"
	"github.com/wildanfaz/backendgolang2_week9/src/libs"
)

type auth_ctrl struct {
	svc interfaces.AuthService
}

func NewCtrl(svc interfaces.AuthService) *auth_ctrl {
	return &auth_ctrl{svc}
}

func (auth *auth_ctrl) SignIn(w http.ResponseWriter, r *http.Request) {
	var data models.User

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		libs.Response(nil, 400, "failed to decode", err).Send(w)
	} else {
		result := auth.svc.Login(data)
		result.Send(w)
	}
}
