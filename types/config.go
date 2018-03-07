package types

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sethdmoore/lemonfeeder/bspwm"
	"github.com/sethdmoore/lemonfeeder/constants"
)

type Config struct {
	MonitorCount int
	Monitors     []string
	Interval     float64
	DesktopCount int
}

func NewConfig() (*Config, error) {
	var c Config
	var err error

	err = envconfig.Process(constants.AppPrefix, &c)
	if err != nil {
		return nil, err
	}

	if c.Interval < 0.0 || c.Interval == 0.0 {
		c.Interval = 1.0
	}

	if c.DesktopCount < 0 || c.DesktopCount == 0 {
		c.DesktopCount = 1
	}

	c.DesktopCount, err = bspwm.GetDesktops()
	if err != nil {
		return nil, err
	}

	c.Monitors, err = bspwm.GetMonitors()
	if err != nil {
		return nil, err
	}

	return &c, nil
}
