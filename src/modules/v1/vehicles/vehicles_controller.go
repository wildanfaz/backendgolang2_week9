package vehicles

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wildanfaz/backendgolang2_week9/src/database/orm/models"
	"github.com/wildanfaz/backendgolang2_week9/src/interfaces"
)

type vehicles_ctrl struct {
	svc interfaces.VehiclesService
}

func NewCtrl(svc interfaces.VehiclesService) *vehicles_ctrl {
	return &vehicles_ctrl{svc}
}

func (ctrl *vehicles_ctrl) GetAllVehicles(w http.ResponseWriter, r *http.Request) {
	data := ctrl.svc.GetAllVehicles()

	if data.IsError != nil {
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *vehicles_ctrl) AddVehicle(w http.ResponseWriter, r *http.Request) {
	var datas models.Vehicle

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {

	} else {
		data := ctrl.svc.AddVehicle(&datas)

		if data.IsError != nil {
			data.Send(w)
		} else {
			data.Send(w)
		}
	}
}

func (ctrl *vehicles_ctrl) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	var datas models.Vehicle

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {

	} else {
		vars := mux.Vars(r)
		data := ctrl.svc.UpdateVehicle(vars["vehicle_id"], &datas)

		if data.IsError != nil {
			data.Send(w)
		} else {
			data.Send(w)
		}
	}
}

func (ctrl *vehicles_ctrl) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	var datas models.Vehicle

	vars := mux.Vars(r)
	data := ctrl.svc.DeleteVehicle(vars["vehicle_id"], &datas)

	if data.IsError != nil {
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *vehicles_ctrl) SearchVehicle(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("vehicle_name")
	data := ctrl.svc.SearchVehicle(search)

	if data.IsError != nil {
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *vehicles_ctrl) PopularVehicles(w http.ResponseWriter, r *http.Request) {
	data := ctrl.svc.PopularVehicles()

	if data.IsError != nil {
		data.Send(w)
	} else {
		data.Send(w)
	}
}
