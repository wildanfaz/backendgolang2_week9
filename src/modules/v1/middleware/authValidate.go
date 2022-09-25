package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/wildanfaz/backendgolang2_week9/src/libs"
)

func CheckAuth(next http.HandlerFunc, role ...[]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			libs.Response(nil, 401, "invalid header", nil).Send(w)
			return
		}

		token := strings.ReplaceAll(headerToken, "Bearer ", "")

		checkToken, err := libs.CheckToken(token)

		if err != nil {
			libs.Response(nil, 401, "invalid token", err).Send(w)
			return
		}

		var checkRole bool
		for _, v := range role[0] {
			if strings.ToLower(v) == strings.ToLower(checkToken.Role) {
				checkRole = true
				break
			}
		}

		if !checkRole {
			libs.Response(nil, 401, "unauthorized role", nil).Send(w)
			return
		}

		ctx := context.WithValue(r.Context(), "name", checkToken.Name)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
