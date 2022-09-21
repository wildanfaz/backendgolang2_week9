package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wildanfaz/backendgolang2_week9/src/database/orm/models"
	"github.com/wildanfaz/backendgolang2_week9/src/interfaces"
)

type users_ctrl struct {
	svc interfaces.UsersService
}

func NewCtrl(svc interfaces.UsersService) *users_ctrl {
	return &users_ctrl{svc}
}

func (ctrl *users_ctrl) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	data := ctrl.svc.GetAllUsers()

	if data.IsError != nil {
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *users_ctrl) AddUser(w http.ResponseWriter, r *http.Request) {
	var datas models.User

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {

	} else {
		data := ctrl.svc.AddUser(&datas)

		if data.IsError != nil {
			data.Send(w)
		} else {
			data.Send(w)
		}
	}
}

func (ctrl *users_ctrl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var datas models.User

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {

	} else {
		vars := mux.Vars(r)
		data := ctrl.svc.UpdateUser(vars["name"], &datas)

		if data.IsError != nil {
			data.Send(w)
		} else {
			data.Send(w)
		}
	}
}

func (ctrl *users_ctrl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var datas models.User

	vars := mux.Vars(r)

	data := ctrl.svc.DeleteUser(vars["name"], &datas)

	if data.IsError != nil {
		data.Send(w)
	} else {
		data.Send(w)
	}
}

// func (re *users_ctrl) SearchUser(w http.ResponseWriter, r *http.Request) {
//
// 	data, err := re.svc.SearchUser(r)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	}

// 	json.NewEncoder(w).Encode(data)
// }
