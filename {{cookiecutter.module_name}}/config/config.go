package config

import (
	"os"
	"{{cookiecutter.module_name}}/logging"

	"gopkg.in/yaml.v3"
)

var logger = logging.GetSugar()
var GlobalConfig Config

type Config struct {
	WebAddr     string `yaml:"web_addr"`
	DatabaseURL string `yaml:"database_url"`
}

func InitConfig(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read config file %s: %v", filename, err)
	}

	err = yaml.Unmarshal(data, &GlobalConfig)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config file %s: %v", filename, err)
	}

	return nil
}
