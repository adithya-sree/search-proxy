package middleware

import (
	"net/http"
	"search/src/app/common"
	"search/src/config"
)

type Middleware struct {
	config *config.Config
}

func NewMiddleware(c *config.Config) Middleware {
	return Middleware{
		config: c,
	}
}

func (m Middleware) AuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, _ := r.BasicAuth()
		if username != m.config.AppConfig.Username || password != m.config.AppConfig.Password {
			common.RespondError(w, http.StatusUnauthorized, "invalid or missing username/password")
			return
		}
		f(w, r)
	}
}
