package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	Port       string        `yaml:"port" default:"8080"`
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
		fileLocation := path + "alir-config.yaml"
		fmt.Println(fileLocation)

	}
	return nil
}
