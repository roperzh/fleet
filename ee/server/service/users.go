package service

import (
	"context"

	"github.com/fleetdm/fleet/v4/server/contexts/ctxerr"
	"github.com/fleetdm/fleet/v4/server/fleet"
	"github.com/fleetdm/fleet/v4/server/ptr"
)

func (svc *Service) GetSSOUser(ctx context.Context, auth fleet.Auth) (*fleet.User, error) {
	config, err := svc.ds.AppConfig(ctx)
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "getting app config")
	}

	if !config.SSOSettings.EnableJITProvisioning {
		return svc.Service.GetSSOUser(ctx, auth)
	}

	u := fleet.UserPayload{
		Email:      ptr.String(auth.UserID()),
		SSOEnabled: ptr.Bool(true),
		GlobalRole: ptr.String(fleet.RoleObserver),
	}

	user, err := svc.NewUser(ctx, u)
	if err != nil {
		return nil, ctxerr.Wrap(ctx, err, "creating new SSO user")
	}

	return user, nil
}
