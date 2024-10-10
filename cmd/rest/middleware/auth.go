package middleware

import (
	"file-server/lib/auth"
	"file-server/lib/config"
	"file-server/lib/helper"
	"log"
	"net/http"
)

type authMiddleware struct {
	apiConfig config.APIConfig
}

func NewAuthMiddleware(apiConfig config.APIConfig) *authMiddleware {
	return &authMiddleware{
		apiConfig: apiConfig,
	}
}

func (m *authMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s [%s] %s", r.RemoteAddr, r.Method, r.URL)
		user, err := auth.ValidateCurrentUser(m.apiConfig, r)
		if err != nil {
			helper.WriteResponse(w, err, nil)
			return
		}

		r = auth.SetUserContext(r, user)
		next.ServeHTTP(w, r)
	})
}
