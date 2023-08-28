package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sulton0011/errs"
)

type (

	// Config
	Config struct {
		App      `json:"app"`
		Postgres `json:"postgres"`
		Logger   `json:"logger"`
		HTTP     `json:"http"`
		GRPC     `json:"grpc"`
		Default  `json:"default"`
	}

	// App -
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}
	// Postgres -
	Postgres struct {
		Url     string ` env-required:"true" yaml:"pg_url" env:"PG_URL"`
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
	}

	// Logger -
	Logger struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// Http -
	HTTP struct {
		Port   string `env-required:"true" yaml:"http_port" env:"HTTP_PORT"`
		Scheme string `env-required:"true" yaml:"http_scheme" env:"HTTP_SCHEME"`
	}

	// gRPC -
	GRPC struct{
		
	}

	// Default -
	Default struct {
		Offset string `env-required:"true" yaml:"offset" env:"OFFSET"`
		Limit  string `env-required:"true" yaml:"limit" env:"LIMIT"`
		Page   string `env-required:"true" yaml:"page" env:"PAGE"`
	}
)

func NewConfig() (cfg *Config, err errs.Error) {
	defer errs.Wrap(&err, "NewConfig")
	cfg = &Config{}

	
	err = cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, errs.Wrap(&err, "cleanenv.ReadConfig")
	}
	
	err = cleanenv.ReadConfig(".env", cfg)
	if err != nil {
		return nil, errs.Wrap(&err, "cleanenv.ReadEnv")
	}

	return cfg, nil
}

