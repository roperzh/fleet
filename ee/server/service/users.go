package service

import (
	"context"
	"errors"

	"github.com/fleetdm/fleet/v4/server/contexts/ctxerr"
	"github.com/fleetdm/fleet/v4/server/fleet"
	"github.com/fleetdm/fleet/v4/server/ptr"
)

func (svc *Service) GetSSOUser(ctx context.Context, auth fleet.Auth) (*fleet.User, error) {
	config, err := svc.ds.AppConfig(ctx)
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "getting app config")
	}

	user, err := svc.Service.GetSSOUser(ctx, auth)
	var nfe fleet.NotFoundError
	if !(errors.As(err, &nfe) && config.SSOSettings.EnableJITProvisioning) {
		return user, err
	}

	user, err = svc.NewUser(ctx, fleet.UserPayload{
		Name:       ptr.String(auth.UserDisplayname()),
		Email:      ptr.String(auth.UserID()),
		SSOEnabled: ptr.Bool(true),
		GlobalRole: ptr.String(fleet.RoleObserver),
	})
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "creating new SSO user")
	}

	return user, nil
}
