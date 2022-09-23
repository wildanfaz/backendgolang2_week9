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

	route.HandleFunc("", middleware.CheckAuth(ctrl.GetAllHistories, []string{"User", "Admin"})).Methods("GET")
	route.HandleFunc("/search", middleware.CheckAuth(ctrl.SearchHistory, []string{"User", "Admin"})).Methods("GET")
	route.HandleFunc("", middleware.CheckAuth(ctrl.AddHistory, []string{"User", "Admin"})).Methods("POST")
	route.HandleFunc("/{history_id}", middleware.CheckAuth(ctrl.UpdateHistory, []string{"Admin"})).Methods("PUT")
	route.HandleFunc("/{history_id}", middleware.CheckAuth(ctrl.DeleteHistory, []string{"Admin"})).Methods("DELETE")
}
