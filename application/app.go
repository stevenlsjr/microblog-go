package application

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"microblog/config"
	"microblog/models"
)

type App struct {
	Config *config.Config
	Db     *gorm.DB
	Users  *UserRepo
}

func New(config *config.Config) (*App, error) {
	db, err := gorm.Open(postgres.Open(config.DatabaseDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	users := &UserRepo{db: db}
	app := &App{Db: db, Config: config, Users: users}
	err = db.AutoMigrate(models.AllModels()...)
	if err != nil {
		log.Error().Err(err).Msg("could not run migrations")
		return nil, err
	}
	return app, nil
}
