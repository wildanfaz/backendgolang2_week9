package auth

import (
	"github.com/wildanfaz/backendgolang2_week9/src/database/orm/models"
	"github.com/wildanfaz/backendgolang2_week9/src/interfaces"
	"github.com/wildanfaz/backendgolang2_week9/src/libs"
)

type auth_service struct {
	repo interfaces.UsersRepo
}

type token_response struct {
	Role  string `json:"role"`
	Token string `json:"token"`
}

func NewService(repo interfaces.UsersRepo) *auth_service {
	return &auth_service{repo}
}

func (auth *auth_service) Login(body models.User) *libs.Resp {
	user, err := auth.repo.FindUserByName(body.Name)

	if err != nil {
		return libs.Response(nil, 401, "incorrect name", err)
	}

	if err := libs.CheckPassword(user.Password, body.Password); err != nil {
		return libs.Response(nil, 401, "incorrect password", err)
	}

	token := libs.NewToken(body.Name, user.Role)
	theToken, err := token.Create()

	if err != nil {
		return libs.Response(nil, 401, "failed create token", err)
	}

	return libs.Response(token_response{Role: user.Role, Token: theToken}, 200, "token created", nil)
}
