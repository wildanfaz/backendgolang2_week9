package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wildanfaz/backendgolang2_week9/src/database/orm/models"
	"github.com/wildanfaz/backendgolang2_week9/src/interfaces"
	"github.com/wildanfaz/backendgolang2_week9/src/libs"
)

type users_ctrl struct {
	svc interfaces.UsersService
}

func NewCtrl(svc interfaces.UsersService) *users_ctrl {
	return &users_ctrl{svc}
}

func (ctrl *users_ctrl) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	claims_users := r.Context().Value("name")
	fmt.Println(claims_users.(string))

	data := ctrl.svc.GetAllUsers()

	if data.IsError != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}

func (ctrl *users_ctrl) GetUser(w http.ResponseWriter, r *http.Request) {
	claims_users := r.Context().Value("name")
	fmt.Println(claims_users.(string))

	data := ctrl.svc.GetUserByName(claims_users.(string))

	if data.IsError != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}

func (ctrl *users_ctrl) AddUser(w http.ResponseWriter, r *http.Request) {
	var datas models.User

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.Response(nil, 400, "failed to decode", err).Send(w)
		return
	}

	data := ctrl.svc.AddUser(&datas)
	if data.IsError != nil {
		data.Send(w)
		return
	}

	data.Send(w)

}

func (ctrl *users_ctrl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var datas models.User

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.Response(nil, 400, "failed to decode", err).Send(w)
		return
	}

	vars := mux.Vars(r)
	data := ctrl.svc.UpdateUser(vars["name"], &datas)

	if data.IsError != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}

func (ctrl *users_ctrl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var datas models.User

	vars := mux.Vars(r)

	data := ctrl.svc.DeleteUser(vars["name"], &datas)

	if data.IsError != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}

func (ctrl *users_ctrl) GetUserByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	data := ctrl.svc.GetUserByName(vars["name"])

	if data.IsError != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}
