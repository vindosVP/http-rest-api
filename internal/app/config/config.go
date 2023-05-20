package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Sever struct {
		BindAddr string `yaml:"bind_addr"`
	} `yaml:"sever"`

	LogLevel string `yaml:"log_level"`

	DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string
		DBName   string `yaml:"db_name"`
	} `yaml:"db"`
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

	dbPwd, ex := os.LookupEnv("DB_PWD")

	if !ex {
		return nil, errors.New("no env variable DB_PWD found")
	} else {
		config.DB.Password = dbPwd
	}

	return config, nil

}
