package mux

import (
	"github.com/HenriquePercu/go-service/apis/services/sales/route/sys/checkapi"
	"github.com/HenriquePercu/go-service/foundation/web"
	"os"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(shutdown chan os.Signal) *web.App {

	mux := web.NewApp(shutdown)

	checkapi.Routes(mux)

	return mux
}
