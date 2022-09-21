package routers

import (
	"errors"

	"github.com/gorilla/mux"
	"github.com/wildanfaz/backendgolang2_week9/src/database/orm"
	"github.com/wildanfaz/backendgolang2_week9/src/modules/v1/histories"
	"github.com/wildanfaz/backendgolang2_week9/src/modules/v1/users"
	"github.com/wildanfaz/backendgolang2_week9/src/modules/v1/vehicles"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	db, err := orm.New()
	if err != nil {
		return nil, errors.New("failed init database")
	}

	users.New(mainRoute, db)
	vehicles.New(mainRoute, db)
	histories.New(mainRoute, db)

	return mainRoute, nil
}
