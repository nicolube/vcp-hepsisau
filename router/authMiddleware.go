package router

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/nicolube/vcp-hepsiau-backend/database"
)

type UserManagerContent string

type UserManager struct {
	Reposetory database.Reposetory
}

func (userManager *UserManager) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			w.WriteHeader(401)
			return
		}
		token := strings.SplitN(authHeader, " ", 2)[1]
		fmt.Println(token)
		tokenModel, err := userManager.Reposetory.GetTokenByToken(token)
		if err != nil {
			w.WriteHeader(401)
			return
		}
		userModel, err := userManager.Reposetory.GetUser(tokenModel.UserId)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		ctx := context.WithValue(r.Context(), UserManagerContent("user"), userModel)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
