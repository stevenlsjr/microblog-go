package api

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"microblog/application"
	"microblog/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var ctx *application.TestSuite = application.NewTestSuite()

func TestMain(m *testing.M) {
	flag.Parse()
	ctx.SetUp()
	defer ctx.TearDown()
	rc := m.Run()

	os.Exit(rc)
}

func TestIndexView(t *testing.T) {
	api := V1(ctx.App)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	api.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 200)
	var bodyJson map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &bodyJson)
	assert.Contains(t, bodyJson, "users")
}

func TestUsers(t *testing.T) {
	api := V1(ctx.App)
	spyDb := ctx.App.Db
	ctx.App.Db = spyDb.Begin()
	t.Cleanup(func() {
		if q := ctx.App.Db.Rollback(); q.Error != nil {
			panic(q.Error)
		}
		ctx.App.Db = spyDb
	})
	var userA models.User
	assert.NoError(t, faker.FakeData(&userA))
	assert.NoError(t, ctx.App.Db.Create(&userA).Error)

	t.Run("Can list users", func(t *testing.T) {

		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/users/", nil)
		assert.NoError(t, err)
		api.ServeHTTP(w, req)
		var bodyData UserList
		assert.Equal(t, 200, w.Code)
		body := w.Body.Bytes()
		assert.NoError(t, json.Unmarshal(body, &bodyData))
		fmt.Printf("")
		assert.Greater(t, int(bodyData.Count), 0)
		assert.Greater(t, len(bodyData.Objects), 0)
	})

	t.Run("Can read user details by id", func(t *testing.T) {
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/users/%s/", userA.ID)
		req, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)
		api.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		var respUser models.User
		assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &respUser))
		assert.Equal(t, userA.ID, respUser.ID)
		assert.Equal(t, userA.Username, respUser.Username)
		assert.Equal(t, userA.CreatedAt.Second(), respUser.CreatedAt.Second())
	})

}
