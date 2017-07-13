package config

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/kelseyhightower/envconfig"
	"github.com/sethdmoore/lemonfeeder/constants"
)

type Config struct {
	monitor_count int
	desktops      int
}

func New() *Config {
	var c Config
	envconfig.Process(constants.AppPrefix, &c)
	spew.Dump(c)
	return &c
}
