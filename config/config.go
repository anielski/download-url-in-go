package config

import (
	"github.com/anielski/download-url-in-go/app"
	"gl.tzv.io/spf13/viper"
	"strings"
)

var (
	appName = `Test API`
	appVersion = "0.1.0"
	devEnv = "dev"
	stagingEnv = "staging"
	prodEnv = "production"
	replacer = strings.NewReplacer(".", "_")
	appConfigPrefix = `APP`
)

// LoadConfig godoc
func LoadConfig() {
	app.API = &app.Application{}
	app.API.Name = appName
	app.API.Version = appVersion
	loadENV(app.API)
	loadAppConfig((app.API))
}

// loadAppConfig: read application config and build viper object
func loadAppConfig(app *app.Application) {
	var (
		appConfig *viper.Viper
	)
	appConfig = viper.New()
	appConfig.SetEnvKeyReplacer(replacer)
	appConfig.SetEnvPrefix(appConfigPrefix)
	appConfig.AutomaticEnv()
	app.Config = *appConfig
}

// loadENV
func loadENV(app *app.Application) {
	var APPENV string
	var appConfig viper.Viper
	appConfig = viper.Viper(app.Config)
	APPENV = appConfig.GetString("ENV")
	switch APPENV {
	case devEnv:
		app.ENV = devEnv
		break
	case stagingEnv:
		app.ENV = stagingEnv
		break
	case prodEnv:
		app.ENV = prodEnv
		break
	default:
		app.ENV = devEnv
		break
	}
}
