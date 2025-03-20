package checkapi

import (
	"github.com/HenriquePercu/go-service/apis/services/api/mid"
	"github.com/HenriquePercu/go-service/app/api/authclient"
	"github.com/HenriquePercu/go-service/business/api/auth"
	"github.com/HenriquePercu/go-service/foundation/logger"
	"github.com/HenriquePercu/go-service/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(app *web.App, log *logger.Logger, authClient *authclient.Client) {
	authen := mid.AuthenticateService(log, authClient)
	athAdminOnly := mid.AuthorizeService(log, authClient, auth.RuleAdminOnly)

	app.HandleFunc("GET /liveness", liveness)
	app.HandleFunc("GET /readiness", readiness)
	app.HandleFunc("GET /testerror", testError)
	app.HandleFunc("GET /testauth", liveness, authen, athAdminOnly)
}
