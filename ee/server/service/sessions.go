package service

import (
	"context"

	"github.com/fleetdm/fleet/v4/server/fleet"
	"github.com/fleetdm/fleet/v4/server/ptr"
)

func (svc *Service) HandleSSOUserMissing(ctx context.Context) error {
	config, err := svc.ds.AppConfig(ctx)
	if err != nil {
		return err
	}

	if !config.SSOSettings.EnableAutomaticProvisioning {
		return svc.Service.HandleSSOUserMissing(ctx)
	}

	u := fleet.UserPayload{
		SSOEnabled: ptr.Bool(true),
		GlobalRole: ptr.String(fleet.RoleAdmin),
	}

	if _, err := svc.CreateUser(ctx, u); err != nil {
		return err
	}

	return nil
}
