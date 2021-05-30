package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	api2 "microblog/api"
	"microblog/application"
	"microblog/config"
)

func main() {

	log.Level(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	err := godotenv.Load(".env")
	if err == nil {
		log.Debug().Msg("loading dotenv")
	}

	cfg, err := config.FromEnv()
	if err != nil {
		panic(fmt.Sprintf("could not load config: %v", err))
	}

	app, err := application.New(cfg)
	if err != nil {
		panic(err)
	}
	api := api2.V1(app)

	if err := api.Run(app.Config.Addr()); err != nil {
		panic(err)
	}

}
