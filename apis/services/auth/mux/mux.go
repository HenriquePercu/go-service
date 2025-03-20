// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"github.com/HenriquePercu/go-service/apis/services/api/mid"
	"github.com/HenriquePercu/go-service/apis/services/auth/route/authapi"
	"github.com/HenriquePercu/go-service/apis/services/auth/route/checkapi"
	"github.com/HenriquePercu/go-service/business/api/auth"
	"github.com/HenriquePercu/go-service/foundation/logger"
	"github.com/HenriquePercu/go-service/foundation/web"
	"os"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(log *logger.Logger, auth *auth.Auth, shutdown chan os.Signal) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics())

	checkapi.Routes(app, auth)
	authapi.Routes(app, auth)

	return app
}
