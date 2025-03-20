package mid

import (
	"context"
	"github.com/HenriquePercu/go-service/app/api/authclient"
	"github.com/HenriquePercu/go-service/app/api/mid"
	"github.com/HenriquePercu/go-service/foundation/logger"
	"github.com/HenriquePercu/go-service/foundation/web"
	"net/http"
)

// AuthorizeService executes the authorize middleware functionality.
func AuthorizeService(log *logger.Logger, client *authclient.Client, rule string) web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			return mid.AuthorizeService(ctx, log, client, rule, hdl)
		}

		return h
	}

	return m
}
