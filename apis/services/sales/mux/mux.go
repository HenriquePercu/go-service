package mux

import (
	"github.com/HenriquePercu/go-service/apis/services/api/mid"
	"github.com/HenriquePercu/go-service/apis/services/sales/route/sys/checkapi"
	"github.com/HenriquePercu/go-service/app/api/authclient"
	"github.com/HenriquePercu/go-service/foundation/logger"
	"github.com/HenriquePercu/go-service/foundation/web"
	"os"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(log *logger.Logger, authClient *authclient.Client, shutdown chan os.Signal) *web.App {

	mux := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics())

	checkapi.Routes(mux, log, authClient)

	return mux
}
