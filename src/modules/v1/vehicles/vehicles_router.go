package vehicles

import (
	"github.com/gorilla/mux"
	"github.com/wildanfaz/backendgolang2_week9/src/modules/v1/middleware"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/api/v1/vehicles").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("", middleware.CheckAuth([]string{"User", "Admin"}, ctrl.GetAllVehicles)).Methods("GET")
	route.HandleFunc("/search", middleware.CheckAuth([]string{"User", "Admin"}, ctrl.SearchVehicle)).Methods("GET")
	route.HandleFunc("/popular", middleware.CheckAuth([]string{"User", "Admin"}, ctrl.PopularVehicles)).Methods("GET")
	route.HandleFunc("", middleware.CheckAuth([]string{"Admin"}, ctrl.AddVehicle)).Methods("POST")
	route.HandleFunc("/{vehicle_id}", middleware.CheckAuth([]string{"Admin"}, ctrl.UpdateVehicle)).Methods("PUT")
	route.HandleFunc("/{vehicle_id}", middleware.CheckAuth([]string{"Admin"}, ctrl.DeleteVehicle)).Methods("DELETE")
}
