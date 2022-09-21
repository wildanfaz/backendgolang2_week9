package interfaces

import (
	"github.com/wildanfaz/backendgolang2_week9/src/database/orm/models"
	"github.com/wildanfaz/backendgolang2_week9/src/libs"
)

type HistoriesRepo interface {
	FindAllHistories() (*models.Histories, error)
	SaveHistory(body *models.History) (*models.History, error)
	ChangeHistory(vars string, body *models.History) (*models.History, error)
	RemoveHistory(vars string, body *models.History) (*models.History, error)
	FindHistory(search string) (*models.Histories, error)
}

type HistoriesService interface {
	GetAllHistories() *libs.Resp
	AddHistory(body *models.History) *libs.Resp
	UpdateHistory(vars string, body *models.History) *libs.Resp
	DeleteHistory(vars string, body *models.History) *libs.Resp
	SearchHistory(search string) *libs.Resp
}
