package config

import (
	"github.com/vindosVp/http-rest-api/internal/app/envconfig"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string
	DBName   string `yaml:"db_name"`
}

type Config struct {
	Sever struct {
		BindAddr string `yaml:"bind_addr"`
	} `yaml:"server"`

	LogLevel string `yaml:"log_level"`

	DB DBConfig `yaml:"db"`

	SessionKey string `yaml:"session_key"`
}

func NewConfig(cPath string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(cPath)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}(file)

	d := yaml.NewDecoder(file)
	if err := d.Decode(config); err != nil {
		return nil, err
	}

	envConf := envConfig.New()
	config.DB.Password = envConf.DBConfig.DBPwd

	return config, nil

}
