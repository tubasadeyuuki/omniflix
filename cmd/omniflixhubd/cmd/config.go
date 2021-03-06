package cmd

import (
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
)

type AppConfig struct {
	serverconfig.Config
}

func initAppConfig() (string, interface{}) {
	srvCfg := serverconfig.DefaultConfig()

	srvCfg.MinGasPrices = "0.001uflix"

	simAppConfig := AppConfig{
		Config: *srvCfg,
	}

	simAppTemplate := serverconfig.DefaultConfigTemplate

	return simAppTemplate, simAppConfig
}
