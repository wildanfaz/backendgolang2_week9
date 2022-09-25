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

	route.HandleFunc("", middleware.HandlerChain(middleware.UserAdmin, middleware.CheckAuth).Then(ctrl.GetAllVehicles)).Methods("GET")
	route.HandleFunc("/search", middleware.HandlerChain(middleware.UserAdmin, middleware.CheckAuth).Then(ctrl.SearchVehicle)).Methods("GET")
	route.HandleFunc("/popular", middleware.HandlerChain(middleware.UserAdmin, middleware.CheckAuth).Then(ctrl.PopularVehicles)).Methods("GET")

	route.HandleFunc("", middleware.HandlerChain(middleware.Admin, middleware.CheckAuth, middleware.UploadFileImage).Then(ctrl.AddVehicle)).Methods("POST")

	route.HandleFunc("/{vehicle_id}", middleware.HandlerChain(middleware.Admin, middleware.CheckAuth).Then(ctrl.UpdateVehicle)).Methods("PUT")
	route.HandleFunc("/{vehicle_id}", middleware.HandlerChain(middleware.Admin, middleware.CheckAuth).Then(ctrl.DeleteVehicle)).Methods("DELETE")

	//**example
	route.HandleFunc("/v", middleware.HandlerChain(middleware.UserAdmin, middleware.CheckAuth, middleware.Hello, middleware.UploadFileImage).Then(ctrl.AddVehicle)).Methods("POST")
}
