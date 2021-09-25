package server

import (
	"fmt"

	"github.com/MonsterYNH/nava/engine"
	"github.com/MonsterYNH/nava/setting"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	app *engine.Engine
)

func init() {
	pflag.String("config_dir", ".", "Configuration folder, the default is the current folder")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	viper.AutomaticEnv()

	var (
		err    error
		config *setting.Config
	)
	config, err = setting.GetConfig(viper.GetString("config_dir"))
	if err != nil {
		panic(err)
	}
	app, err = engine.New(config)
	if err != nil {
		panic(err)
	}
}

func Run() error {
	return engine.Run(app)
}

func RegisterAPIV1Service(service engine.Service) error {
	return app.RegisterAPIService(fmt.Sprintf("/v1/%s", service.Name()), service)
}
