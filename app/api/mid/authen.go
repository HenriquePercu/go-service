package mid

import (
	"context"
	"fmt"
	"github.com/HenriquePercu/go-service/app/api/authclient"
	"github.com/HenriquePercu/go-service/app/api/errs"
	"github.com/HenriquePercu/go-service/business/api/auth"
	"github.com/HenriquePercu/go-service/foundation/logger"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"strings"
	"time"
)

// AuthenticateService validates authentication via the auth service.
func AuthenticateService(ctx context.Context, log *logger.Logger, client *authclient.Client, authorization string, handler Handler) error {
	resp, err := client.Authenticate(ctx, authorization)
	if err != nil {
		return errs.New(errs.Unauthenticated, err)
	}

	ctx = setUserID(ctx, resp.UserID)
	ctx = setClaims(ctx, resp.Claims)

	return handler(ctx)
}

// AuthenticateLocal processes the authentication requirements locally.
func AuthenticateLocal(ctx context.Context, auth *auth.Auth, authorization string, handler Handler) error {
	var err error
	parts := strings.Split(authorization, " ")

	switch parts[0] {
	case "Bearer":
		ctx, err = processJWT(ctx, auth, authorization)
	case "Basic":
		ctx, err = processBasic(ctx)
	}

	if err != nil {
		return err
	}

	return handler(ctx)
}

func processBasic(ctx context.Context) (context.Context, error) {
	claims := auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "38dc9d84-018b-4a15-b958-0b78af11c301",
			Issuer:    "service project",
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(8760 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
		Roles: []string{"ADMIN"},
	}

	subjectID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return ctx, errs.Newf(errs.Unauthenticated, "parsing subject: %s", err)
	}

	ctx = setUserID(ctx, subjectID)
	ctx = setClaims(ctx, claims)

	return ctx, nil
}

func processJWT(ctx context.Context, auth *auth.Auth, token string) (context.Context, error) {
	claims, err := auth.Authenticate(ctx, token)
	if err != nil {
		return ctx, errs.New(errs.Unauthenticated, err)
	}

	if claims.Subject == "" {
		return ctx, errs.Newf(errs.Unauthenticated, "authorize: you are not authorized for that action, no claims")
	}

	subjectID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return ctx, errs.New(errs.Unauthenticated, fmt.Errorf("parsing subject: %w", err))
	}

	ctx = setUserID(ctx, subjectID)
	ctx = setClaims(ctx, claims)

	return ctx, nil
}
