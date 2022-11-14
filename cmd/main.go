package main

import (
	"github.com/kelseyhightower/envconfig"
	"tech-user-balance/config"
	"tech-user-balance/infrastructure/repository"
	"tech-user-balance/pkg"
)

func main() {
	log := pkg.GetLogger()

	var cnf config.Config
	err := envconfig.Process("", &cnf)
	if err != nil {
		log.Fatal(err)
		return
	}

	db := repository.ConnectDB(&cnf, log)
}
