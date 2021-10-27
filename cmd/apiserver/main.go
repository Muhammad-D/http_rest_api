package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Muhammad-D/http_rest_api/internal/app/apiserver"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "./configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	//Configuration settings...
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	//API Server settings...
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
