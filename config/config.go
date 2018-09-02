package config

import (
	"io/ioutil"

	"github.com/alexandrevilain/monit/job"
	yaml "gopkg.in/yaml.v2"
)

type HttpConfig struct {
	WorkingResponse string `yaml:"workingResponse"`
	ErrorResponse   string `yaml:"errorResponse"`
	ListenAddress   string `yaml:"listenAddr"`
}

// Config is the struct corressponding to the config of the app
type Config struct {
	Services []job.Service
	HTTP     HttpConfig `yaml:"http"`
}

// LoadConfigFromFile reads the yaml file provided in parameters
// and returns Config struct from it
func LoadConfigFromFile(path string) (Config, error) {
	var cfg Config
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	err = yaml.Unmarshal(fileContent, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
