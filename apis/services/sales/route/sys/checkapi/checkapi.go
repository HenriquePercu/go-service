package checkapi

import (
	"context"
	"github.com/HenriquePercu/go-service/app/api/errs"
	"github.com/HenriquePercu/go-service/foundation/web"
	"math/rand"
	"net/http"
)

func liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{Status: "OK"}

	return web.Respond(ctx, w, status, http.StatusOK)
}

func readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{Status: "OK"}

	return web.Respond(ctx, w, status, http.StatusOK)
}

func testError(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		return errs.Newf(errs.FailedPrecondition, "this is an error")
	}

	status := struct {
		Status string
	}{Status: "OK"}

	return web.Respond(ctx, w, status, http.StatusOK)
}
