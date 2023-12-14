package config

import (
	u "exchangerate/utils"
)

var App *app

type appEnv string

const (
	APP_LOCAL appEnv = "local"
	APP_PROD  appEnv = "production"
	APP_STG   appEnv = "staging"
)

type app struct {
	Env   appEnv
	Debug bool
}

func init() {
	App = &app{
		Env:   appEnv(u.Env("APP_ENV", string(APP_LOCAL))),
		Debug: u.EnvBool("APP_DEBUG", true),
	}
}

func (cfg *app) IsLocal() bool {
	return cfg.Env == APP_LOCAL
}

func (cfg *app) IsProduction() bool {
	return cfg.Env == APP_PROD
}

func (cfg *app) IsStaging() bool {
	return cfg.Env == APP_STG
}
