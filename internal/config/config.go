package config

import (
	"encoding/json"
	"github.com/ilyakaznacheev/cleanenv"
	"goEdu/pkg/logging"
	"os"
	"sync"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Storage `yaml:"storage"`
}

type Storage struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

var once sync.Once
var instance *Config

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		confPath := "config.yaml"
		if value, ok := os.LookupEnv("CONFIG_FILE"); ok {
			confPath = value
		} else {
			logger.Info("env: CONFIG_FILE not set. Default: \"config.yaml\"")
		}

		logger.Debugf("Try to read config file %s", confPath)
		instance = &Config{}
		err := cleanenv.ReadConfig(confPath, instance)
		if err != nil {
			logger.Fatal("Failed to read config file. Error", err.Error())
		}
		configJSON, _ := json.Marshal(&instance)
		logger.Debug("Config file read. Config: ", string(configJSON))
	})
	return instance
}
