package interfaces

import (
	"github.com/wildanfaz/backendgolang2_week9/src/database/orm/models"
	"github.com/wildanfaz/backendgolang2_week9/src/libs"
)

type UsersRepo interface {
	FindAllUsers() (*models.Users, error)
	FindUserByName(name string) (*models.User, error)
	SaveUser(body *models.User) (*models.User, error)
	ChangeUser(vars string, body *models.User) (*models.User, error)
	RemoveUser(vars string, body *models.User) (*models.User, error)
}

type UsersService interface {
	GetAllUsers() *libs.Resp
	GetUserByName(name string) *libs.Resp
	AddUser(body *models.User) *libs.Resp
	UpdateUser(vars string, body *models.User) *libs.Resp
	DeleteUser(vars string, body *models.User) *libs.Resp
}
