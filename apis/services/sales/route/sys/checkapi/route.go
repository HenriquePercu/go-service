package checkapi

import (
	"github.com/HenriquePercu/go-service/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(app *web.App) {
	app.HandleFunc("GET /liveness", liveness)
	app.HandleFunc("GET /readiness", readiness)
}