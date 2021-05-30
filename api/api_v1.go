package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"microblog/application"
)

type ErrorBody struct {
	Message string `json:"message"`
	Status  string `json:"message"`
}

func V1(app *application.App) *gin.Engine {
	router := gin.Default()
	userV1 := NewUsersV1(app)

	router.GET("/", func(c *gin.Context) {
		log.Debug().Msg(c.Request.Host)
		index := map[string]string{
			"users": "/users/",
		}
		c.JSON(200, index)
	})

	userRoutes := router.Group("/users")
	userRoutes.GET("/", userV1.List)
	userRoutes.GET("/:id/", userV1.Detail)

	userRoutes.POST("/", userV1.Create)
	userRoutes.PUT("/:id/", userV1.Update)
	userRoutes.PATCH("/:id/", userV1.PartialUpdate)
	userRoutes.DELETE("/:id/", userV1.Delete)

	return router
}
