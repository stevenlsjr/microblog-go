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
}

func New(config *config.Config) (*App, error) {
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	}
	dsn := postgres.Open(config.DatabaseDSN)
	db, err := gorm.Open(dsn, gormConfig)
	if err != nil {
		return nil, err
	}

	app := &App{Db: db, Config: config}

	if err := app.Init(); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Init() error {
	db := a.Db

	if q := db.Exec("CREATE EXTENSION IF NOT EXISTS pgcrypto"); q.Error != nil {
		return q.Error
	}
	if err := db.AutoMigrate(models.AllModels()...); err != nil {
		log.Error().Err(err).Msg("could not run migrations")
		return err
	}
	return nil
}

func FromEnv() (*App, error) {
	cfg, err := config.FromEnv()
	if err != nil {
		return nil, err
	}
	return New(cfg)
}
