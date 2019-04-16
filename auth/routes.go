package auth

import (
	"jwt-tools-online-shit/config"

	"github.com/go-chi/chi"

	authorization "github.com/travelgateX/go-jwt-tools"
	"github.com/travelgateX/go-jwt-tools/jwt"
)

// Routes : token routes
func Routes(c config.AppConfig) *chi.Mux {
	r := chi.NewRouter()

	// jwt-tools middleware
	jwtParser := jwt.NewParser(c.GetJwtParserConfig())
	authMiddleware := authorization.Middleware(jwtParser)
	r.Use(authMiddleware)

	// token routes
	r.Get("/validgroups", GetValidGroups)
	r.Get("/haspermission", HasPermission)

	return r
}
