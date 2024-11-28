package mid

import (
	"context"
	"fmt"
	"github.com/HenriquePercu/go-service/foundation/logger"
	"time"
)

// Logger writes information about the request to the logs.
func Logger(ctx context.Context, log *logger.Logger, path string, rawQuery string, method string, remoteAddr string, handler Handler) error {
	now := time.Now()

	if rawQuery != "" {
		path = fmt.Sprintf("%s?%s", path, rawQuery)
	}

	log.Info(ctx, "request started", "method", method, "path", path, "remoteaddr", remoteAddr)

	err := handler(ctx)

	log.Info(ctx, "request completed", "method", method, "path", path, "remoteaddr", remoteAddr,
		"statuscode", "since", time.Since(now).String())

	return err
}
