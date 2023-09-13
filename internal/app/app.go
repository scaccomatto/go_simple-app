package app

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"paru.net/gosimpleapp/internal/configuration"
	"paru.net/gosimpleapp/internal/database"
	"paru.net/gosimpleapp/internal/service"
)

func Start() error {

	if err := configuration.LoadAppConfig(); err != nil {
		log.Error().Err(err).Msg("no environment config found, using default config")
		return err
	}

	// Setting up services
	db := database.NewMapDb()

	fooSvc := service.NewFooService(db)
	fooSvc2 := service.NewFooService(db)
	fooSvc.InsertFoo("gianni", 1)
	fooSvc.InsertFoo("pinotto", 2)
	fooSvc.InsertFoo("pinco", 3)

	fmt.Println(fooSvc2.GetFoo("1"))
	fmt.Println(fooSvc2.GetFoo("2"))
	fmt.Println(fooSvc2.GetFoo("3"))

	return nil

}
