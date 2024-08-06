package main

import (
	"log"

	"github.com/MaximDik/http-rest-api/internal/app/apiserver"
)

func main() {
	s := apiserver.NewAPIServer()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
