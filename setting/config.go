package setting

import "github.com/spf13/viper"

func GetConfig(dir string) (*Config, error) {
	viper.AddConfigPath(dir)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	config := Config{}

	return &config, viper.Unmarshal(&config)
}

type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	DataSource DataSourceConfig `mapstructure:"data_source"`
	Logger     LoggerConfig     `mapstructure:"logger"`
}

type ServerConfig struct {
	Host              string `mapstructure:"host"`
	Port              int    `mapstructure:"port"`
	EnableSwagger     bool   `mapstructure:"enable_swagger"`
	EnablePPROF       bool   `mapstructure:"enable_pprof"`
	EnableAuthCheck   bool   `mapstructure:"enable_auth"`
	EnableHealthCheck bool   `mapstructure:"enable_health"`
}

type DataSourceConfig struct {
	Postgres  PostgresConfig   `mapstructure:"postgres"`
	Promethes PrometheusConfig `mapstructure:"prometheus"`
}

type PostgresConfig struct {
	Enable bool   `mapstructure:"enable"`
	URI    string `mapstructure:"uri"`
}

type PrometheusConfig struct {
	Enable bool   `mapstructure:"enable"`
	URL    string `mapstructure:"url"`
}

type LoggerConfig struct {
	LogLevel        string `mapstructure:"log_level"`
	LogFileName     string `mapstructure:"log_filename"`
	LogMaxAge       int    `mapstructure:"log_max_age"`
	LogRotationTime int    `mapstructure:"log_rotation_time"`
}
