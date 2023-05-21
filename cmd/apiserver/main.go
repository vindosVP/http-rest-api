package main

import (
	"flag"
	"github.com/vindosVp/http-rest-api/internal/app/apiserver"
	"github.com/vindosVp/http-rest-api/internal/app/config"
	"github.com/vindosVp/http-rest-api/internal/app/logger"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.yaml", "path to config file")
}

func main() {

	flag.Parse()

	cfg, err := config.NewConfig(configPath)

	if err != nil {
		log.Fatal(err)
	}

	if err := logger.ConfigureLogger(cfg.LogLevel); err != nil {
		panic("logger configuring failed")
	}

	err = apiserver.Start(cfg)

	if err != nil {
		log.Fatal(err)
	}

}
