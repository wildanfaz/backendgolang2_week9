package interfaces

import (
	"github.com/wildanfaz/backendgolang2_week9/src/database/orm/models"
	"github.com/wildanfaz/backendgolang2_week9/src/libs"
)

type VehiclesRepo interface {
	FindAllVehicles() (*models.Vehicles, error)
	SaveVehicle(body *models.Vehicle) (*models.Vehicle, error)
	ChangeVehicle(vars string, body *models.Vehicle) (*models.Vehicle, error)
	RemoveVehicle(vars string, body *models.Vehicle) (*models.Vehicle, error)
	FindVehicle(search string) (*models.Vehicles, error)
	RatingVehicles() (*models.Vehicles, error)
}

type VehiclesService interface {
	GetAllVehicles() *libs.Resp
	AddVehicle(body *models.Vehicle) *libs.Resp
	UpdateVehicle(vars string, body *models.Vehicle) *libs.Resp
	DeleteVehicle(vars string, body *models.Vehicle) *libs.Resp
	SearchVehicle(search string) *libs.Resp
	PopularVehicles() *libs.Resp
}
