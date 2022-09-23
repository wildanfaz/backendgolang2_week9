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

	route.HandleFunc("", middleware.CheckAuth(ctrl.GetAllVehicles, []string{"User", "Admin"})).Methods("GET")
	route.HandleFunc("/search", middleware.CheckAuth(ctrl.SearchVehicle, []string{"User", "Admin"})).Methods("GET")
	route.HandleFunc("/popular", middleware.CheckAuth(ctrl.PopularVehicles, []string{"User", "Admin"})).Methods("GET")

	route.HandleFunc("", middleware.CheckAuth(middleware.UploadFile(ctrl.AddVehicle), []string{"Admin"})).Methods("POST")

	route.HandleFunc("/{vehicle_id}", middleware.CheckAuth(ctrl.UpdateVehicle, []string{"Admin"})).Methods("PUT")
	route.HandleFunc("/{vehicle_id}", middleware.CheckAuth(ctrl.DeleteVehicle, []string{"Admin"})).Methods("DELETE")

	//**example
	route.HandleFunc("/v", middleware.HandlerChain(middleware.Hello, middleware.UploadFile).Then(ctrl.AddVehicle)).Methods("POST")
}
