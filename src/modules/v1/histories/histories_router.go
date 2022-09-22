package histories

import (
	"github.com/gorilla/mux"
	"github.com/wildanfaz/backendgolang2_week9/src/modules/v1/middleware"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/api/v1/histories").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("", middleware.CheckAuth([]string{"User", "Admin"}, ctrl.GetAllHistories)).Methods("GET")
	route.HandleFunc("/search", middleware.CheckAuth([]string{"User", "Admin"}, ctrl.SearchHistory)).Methods("GET")
	route.HandleFunc("", middleware.CheckAuth([]string{"User", "Admin"}, ctrl.AddHistory)).Methods("POST")
	route.HandleFunc("/{history_id}", middleware.CheckAuth([]string{"Admin"}, ctrl.UpdateHistory)).Methods("PUT")
	route.HandleFunc("/{history_id}", middleware.CheckAuth([]string{"Admin"}, ctrl.DeleteHistory)).Methods("DELETE")
}
