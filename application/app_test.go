package application

import (
	"flag"
	"github.com/jackc/pgconn"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"microblog/models"
	"os"
	"testing"
)

var testContext TestSuite

func TestMain(m *testing.M) {
	testContext.SetUp()
	flag.Parse()
	res := m.Run()
	testContext.TearDown()
	os.Exit(res)
}

func TestUsers(t *testing.T) {
	db := testContext.App.Db.Begin()
	var count int64
	testContext.App.Db.Model(&models.User{}).Count(&count)
	log.Info().Msgf("count for users is now %d", count)
	t.Cleanup(func() {
		log.Info().Msg("Rollback")

		db.Rollback()
	})

	t.Run("Can be created", func(t *testing.T) {
		user := models.User{
			Username: "joebob",
		}
		q := db.Create(&user)
		assert.Nil(t, q.Error)

		var newUser models.User
		q = db.Where("username = ?", "joebob").
			Find(&newUser)
		assert.Nil(t, q.Error)
		assert.Equal(t, "joebob", newUser.Username)
		assert.Equal(t, user.UuidModel.ID, newUser.UuidModel.ID)
		assert.Equal(t, user.UuidModel.UpdatedAt.Second(), newUser.UuidModel.UpdatedAt.Second())

	})

	t.Run("Can be created with an explicit uuid", func(t *testing.T) {
		id := "649ddb55-7d76-4e4c-814b-0e8073cc8f4a"
		user := models.User{
			UuidModel: models.UuidModel{
				ID: id,
			},
			Username: "joebob2",
		}
		assert.Nil(t, db.Create(&user).Error)

		assert.Equal(t, user.ID, id)
		assert.Equal(t, user.Username, "joebob2")
	})

	t.Run("Raises an integrity error if User already exists", func(t *testing.T) {
		assert.Nil(t, db.Create(&models.User{
			Username: "joebob3",
		}).Error)
		err := db.Create(&models.User{
			Username: "joebob3",
		}).Error
		assert.NotNil(t, err)
		assert.IsType(t, err, &pgconn.PgError{})
	})

}
