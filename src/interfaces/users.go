package interfaces

import (
	"github.com/wildanfaz/backendgolang2_week9/src/database/orm/models"
	"github.com/wildanfaz/backendgolang2_week9/src/libs"
)

type UsersRepo interface {
	FindAllUsers() (*models.Users, error)
	SaveUser(body *models.User) (*models.User, error)
	ChangeUser(vars string, body *models.User) (*models.User, error)
	RemoveUser(vars string, body *models.User) (*models.User, error)
	// FindUser(r *http.Request) (*models.Users, error)
}

type UsersService interface {
	GetAllUsers() *libs.Resp
	AddUser(body *models.User) *libs.Resp
	UpdateUser(vars string, body *models.User) *libs.Resp
	DeleteUser(vars string, body *models.User) *libs.Resp
	// SearchUser(r *http.Request) (*models.Users, error)
}
