package main

import (
	_ "github.com/SAIKAII/skResk-Account/apis/web"
	_ "github.com/SAIKAII/skResk-Account/core/accounts"
	"github.com/SAIKAII/skResk-Infra"
	"github.com/SAIKAII/skResk-Infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.HookStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.IrisServerStarter{})
}
