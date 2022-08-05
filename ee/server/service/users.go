package service

import (
	"context"

	"github.com/fleetdm/fleet/v4/server/contexts/ctxerr"
	"github.com/fleetdm/fleet/v4/server/fleet"
	"github.com/fleetdm/fleet/v4/server/ptr"
)

func (svc *Service) HandleSSOUserMissing(ctx context.Context, email string) error {
	config, err := svc.ds.AppConfig(ctx)
	if err != nil {
		return ctxerr.Wrap(ctx, err, "getting app config")
	}

	if !config.SSOSettings.EnableJITProvisioning {
		return svc.Service.HandleSSOUserMissing(ctx, email)
	}

	u := fleet.UserPayload{
		Email:      &email,
		SSOEnabled: ptr.Bool(true),
		GlobalRole: ptr.String(fleet.RoleObserver),
	}

	if _, err := svc.Service.NewUser(ctx, u); err != nil {
		return ctxerr.Wrap(ctx, err, "creating new SSO user")
	}

	return nil
}
