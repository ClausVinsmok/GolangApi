package apiserver

import "golang-rest-api/internel/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: "",
		LogLevel: "",
		Store:    store.NewConfig(),
	}
}
