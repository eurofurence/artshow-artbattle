package configuration

import (
	"github.com/eurofurence/artshow-artbattle/internal/application/middleware"
	"github.com/eurofurence/artshow-artbattle/internal/application/server"
	"github.com/eurofurence/artshow-artbattle/internal/repository/idp"
	"github.com/eurofurence/artshow-artbattle/internal/repository/logging"
	"github.com/eurofurence/artshow-artbattle/internal/repository/vault"
	auconfigapi "github.com/StephanHCB/go-autumn-config-api"
	auconfigenv "github.com/StephanHCB/go-autumn-config-env"
	aulogging "github.com/StephanHCB/go-autumn-logging"
)

func Setup() error {
	if err := auconfigenv.Setup(ConfigItems(), warn); err != nil {
		return err
	}
	if err := auconfigenv.Read(); err != nil {
		return err
	}
	if err := auconfigenv.Validate(); err != nil {
		return err
	}
	return nil
}

func ConfigItems() []auconfigapi.ConfigItem {
	return join(
		logging.ConfigItems(),
		server.ConfigItems(),
		middleware.CorsConfigItems(),
		middleware.SecurityConfigItems(),
		vault.ConfigItems(),
		idp.ConfigItems(),
		// add new config item providers here
	)
}

func join(configs ...[]auconfigapi.ConfigItem) []auconfigapi.ConfigItem {
	result := make([]auconfigapi.ConfigItem, 0)
	for _, items := range configs {
		result = append(result, items...)
	}
	return result
}

func warn(message string) {
	aulogging.Logger.NoCtx().Warn().Print(message)
}
