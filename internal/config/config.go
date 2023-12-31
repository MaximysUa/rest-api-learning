package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"rest-api-learning/pkg/logging"
	"sync"
)

// создаём структуру аналагичную ямл файлу
type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  struct {
		Type   string `yaml:"type"`
		BindIP string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"listen"`

	Storage StorageConfig `yaml:"storage"`
}
type StorageConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Создание синглтона
var instance *Config
var once sync.Once

func GetConfig() *Config {
	//код ниже выполниться только 1 раз
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
