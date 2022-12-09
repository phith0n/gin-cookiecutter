package config

import (
	"os"
	"{{cookiecutter.module_name}}/logging"
)

var logger = logging.GetSugar()
var GlobalConfig Config

type Config struct {
	WebAddr string `yaml:"web_addr"`
}

func InitConfig(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		logger.Errorf("failed to read config file %s: %v", filename, err)
		return err
	}

	err = yaml.Unmarshal(data, &GlobalConfig)
	if err != nil {
		logger.Errorf("failed to unmarshal config file %s: %v", filename, err)
		return err
	}

	return nil
}
