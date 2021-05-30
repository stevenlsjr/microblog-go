package application

import (
	"github.com/rs/zerolog/log"
)

type CreateTestDbOptions struct {
	createDb bool
}

func (a *App) CreateTestApp(options *CreateTestDbOptions) (*App, error) {
	newConfig, err := a.Config.TestConfig()
	if err != nil {
		return nil, err
	}
	if options.createDb {
		if err = a.createTestDbIfNotExists(newConfig.DbName); err != nil {
			return nil, err
		}
	}
	return New(&newConfig.Config)
}

type countResult struct {
	Count int
}

func (a *App) createTestDbIfNotExists(dbName string) (err error) {
	log.Info().Msgf("Creating table '%s' if it exists", dbName)

	db := a.Db
	var result countResult
	if q := db.Raw(`
				SELECT count(1) FROM pg_catalog.pg_database
				WHERE datname = ?
			`, dbName).Scan(&result); q.Error != nil {
		return q.Error
	}

	if result.Count == 0 {

		if q := db.Exec("CREATE DATABASE " + dbName); q.Error != nil {
			return q.Error
		}
	}
	return nil
}

func (a *App) dropTestDbIfExists(databaseName string) (err error) {
	db := a.Db
	var result countResult
	if q := db.Raw(`
				SELECT count(1) FROM pg_catalog.pg_database
				WHERE datname = ?
			`, databaseName).Scan(&result); q.Error != nil {
		return q.Error
	}

	if result.Count != 0 {

		if q := db.Exec("DROP DATABASE " + databaseName); q.Error != nil {
			return q.Error
		}
	}
	return nil
}
