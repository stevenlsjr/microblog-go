package application

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"microblog/config"
)

type TestSuite struct {
	testOptions CreateTestDbOptions
	NonTestApp  *App
	App         *App
	isSetUp     bool
}

func NewTestSuite() *TestSuite {

	suite := TestSuite{testOptions: CreateTestDbOptions{createDb: true}, isSetUp: false}
	return &suite
}

func (t *TestSuite) SetUp() {
	if t.isSetUp {
		panic("Can't re-setup test suite")
	}
	err := godotenv.Load(".env", "../.env")
	if err != nil {
		log.Error().Err(err)
	}
	if t.NonTestApp, err = FromEnv(); err != nil {
		panic(err)
	}
	t.testOptions = CreateTestDbOptions{createDb: true}

	if t.App, err = t.NonTestApp.CreateTestApp(&t.testOptions); err != nil {
		panic(err)
	}
	t.isSetUp = true
}

func (t *TestSuite) TearDown() {
	var err error
	log.Info().Msg("Tearing down tests")
	testDbName, err := config.GetTestDbName(t.App.Config.DatabaseDSN)
	if err != nil {
		panic(err)
	}
	if err = t.NonTestApp.dropTestDbIfExists(testDbName.DbName); err != nil {
		panic(err)
	}
}
