package web

import (
	"context"
	"time"
)

type ctxKey int

const (
	key ctxKey = 1
)

type Values struct {
	TraceID    string
	Now        time.Time
	StatusCode int
}

func GetValues(ctx context.Context) *Values {
	v, ok := ctx.Value(key).(*Values)

	if !ok {
		return &Values{
			TraceID: "000000",
			Now:     time.Now(),
		}
	}

	return v
}
func setStatusCode(ctx context.Context, statusCode int) {
	v, ok := ctx.Value(key).(*Values)

	if !ok {
		return
	}

	v.StatusCode = statusCode
}

func setValues(ctx context.Context, values Values) context.Context {
	return context.WithValue(ctx, key, values)
}
