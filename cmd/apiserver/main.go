package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/MaximDik/http-rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {

	flag.Parse()

	config := apiserver.NewConfig()
	fmt.Printf("Config path before decode file %+v\n", config)

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Config path after decode file %+v\n", config)
	s := apiserver.NewAPIServer(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
