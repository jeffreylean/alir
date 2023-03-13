package config

import (
	"os"
	"time"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	Port       int32         `yaml:"port" default:"8080"`
	Timeout    time.Duration `yaml:"timeout" default:"60s"`
	ConfigFile string
}

func (c *Config) Load() error {
	if c.ConfigFile == "" {
		path, err := os.Getwd()
		if err != nil {
			return err
		}

		// alir-config.yaml is the default file name
		c.ConfigFile = path + "/alir-config.yaml"
	}

	yFile, err := os.ReadFile(c.ConfigFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yFile, &c)
	if err != nil {
		return err
	}

	return nil
}
