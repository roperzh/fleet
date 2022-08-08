package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/fleetdm/fleet/v4/server/contexts/ctxerr"
	"github.com/fleetdm/fleet/v4/server/fleet"
	"github.com/fleetdm/fleet/v4/server/ptr"
)

func (svc *Service) GetSSOUser(ctx context.Context, auth fleet.Auth) (*fleet.User, error) {
	config, err := svc.ds.AppConfig(ctx)
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "getting app config")
	}

	user, err := svc.ds.UserByEmail(ctx, auth.UserID())
	var nfe fleet.NotFoundError
	fmt.Println("=----------------------------------", errors.As(err, &nfe))
	if errors.As(err, &nfe) && !config.SSOSettings.EnableJITProvisioning {
		return nil, err
	}

	user, err = svc.NewUser(ctx, fleet.UserPayload{
		Name:       ptr.String(""),
		Email:      ptr.String(auth.UserID()),
		SSOEnabled: ptr.Bool(true),
		GlobalRole: ptr.String(fleet.RoleObserver),
	})
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "creating new SSO user")
	}

	return user, nil
}
