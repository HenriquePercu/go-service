package mid

import (
	"context"
	"errors"
	"github.com/HenriquePercu/go-service/app/api/authclient"
	"github.com/HenriquePercu/go-service/app/api/errs"
	"github.com/HenriquePercu/go-service/foundation/logger"
)

// ErrInvalidID represents a condition where the id is not a uuid.
var ErrInvalidID = errors.New("ID is not in its proper form")

// AuthorizeService executes the specified role and does not extract any domain data.
func AuthorizeService(ctx context.Context, log *logger.Logger, client *authclient.Client, rule string, handler Handler) error {
	userID, err := GetUserID(ctx)
	if err != nil {
		return errs.New(errs.Unauthenticated, err)
	}

	auth := authclient.Authorize{
		Claims: GetClaims(ctx),
		UserID: userID,
		Rule:   rule,
	}

	if err := client.Authorize(ctx, auth); err != nil {
		return errs.New(errs.Unauthenticated, err)
	}

	return handler(ctx)
}
