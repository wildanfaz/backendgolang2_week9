package histories

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wildanfaz/backendgolang2_week9/src/database/orm/models"
	"github.com/wildanfaz/backendgolang2_week9/src/interfaces"
)

type histories_ctrl struct {
	svc interfaces.HistoriesService
}

func NewCtrl(svc interfaces.HistoriesService) *histories_ctrl {
	return &histories_ctrl{svc}
}

func (ctrl *histories_ctrl) GetAllHistories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := ctrl.svc.GetAllHistories()

	if data.IsError != nil {
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *histories_ctrl) AddHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.History

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {

	} else {
		data := ctrl.svc.AddHistory(&datas)

		if data.IsError != nil {
			data.Send(w)
		} else {
			data.Send(w)
		}
	}
}

func (ctrl *histories_ctrl) UpdateHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.History

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {

	} else {
		vars := mux.Vars(r)
		data := ctrl.svc.UpdateHistory(vars["history_id"], &datas)

		if data.IsError != nil {
			data.Send(w)
		} else {
			data.Send(w)
		}
	}
}

func (ctrl *histories_ctrl) DeleteHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var datas models.History

	vars := mux.Vars(r)
	data := ctrl.svc.DeleteHistory(vars["history_id"], &datas)

	if data.IsError != nil {
		data.Send(w)
	} else {
		data.Send(w)
	}
}

func (ctrl *histories_ctrl) SearchHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	search := r.URL.Query().Get("vehicle_id")
	data := ctrl.svc.SearchHistory(search)

	if data.IsError != nil {
		data.Send(w)
	} else {
		data.Send(w)
	}
}
