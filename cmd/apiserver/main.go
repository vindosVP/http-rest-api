package main

import (
	"flag"
	"github.com/vindosVp/http-rest-api/internal/app/apiserver"
	"github.com/vindosVp/http-rest-api/internal/app/config"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.yaml", "path to config file")
}

func main() {

	flag.Parse()

	cfg, err := config.NewConfig(configPath)

	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(cfg)

	if err := s.Start(); err != nil {
		log.Fatalln(err)
	}

}
