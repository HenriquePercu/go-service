package mid

import (
	"context"
	"github.com/HenriquePercu/go-service/app/api/authclient"
	"github.com/HenriquePercu/go-service/app/api/mid"
	"github.com/HenriquePercu/go-service/business/api/auth"
	"github.com/HenriquePercu/go-service/foundation/logger"
	"github.com/HenriquePercu/go-service/foundation/web"
	"net/http"
)

// AuthenticateService validates authentication via the auth service.
func AuthenticateService(log *logger.Logger, client *authclient.Client) web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			return mid.AuthenticateService(ctx, log, client, r.Header.Get("authorization"), hdl)
		}

		return h
	}

	return m
}

// AuthenticateLocal processes the authentication requirements locally.
func AuthenticateLocal(auth *auth.Auth) web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}

			return mid.AuthenticateLocal(ctx, auth, r.Header.Get("authorization"), hdl)
		}

		return h
	}

	return m
}
