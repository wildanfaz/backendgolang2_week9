package users

import (
	"github.com/gorilla/mux"
	"github.com/wildanfaz/backendgolang2_week9/src/modules/v1/middleware"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/api/v1/users").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	//** unused get all users
	// route.HandleFunc("", middleware.CheckAuth([]string{"Admin"}, ctrl.GetAllUsers)).Methods("GET")
	route.HandleFunc("", middleware.HandlerChain(middleware.UserAdmin, middleware.CheckAuth).Then(ctrl.GetUser)).Methods("GET")
	route.HandleFunc("/{name}", middleware.HandlerChain(middleware.UserAdmin, middleware.CheckAuth).Then(ctrl.GetUserByName)).Methods("GET")

	//** register
	route.HandleFunc("", ctrl.AddUser).Methods("POST")

	route.HandleFunc("/{name}", middleware.HandlerChain(middleware.Admin, middleware.CheckAuth).Then(ctrl.UpdateUser)).Methods("PUT")
	route.HandleFunc("/{name}", middleware.HandlerChain(middleware.Admin, middleware.CheckAuth).Then(ctrl.DeleteUser)).Methods("DELETE")
}
