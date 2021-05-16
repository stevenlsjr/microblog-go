package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"microblog/application"
)

type ErrorBody struct {
	Message string `json: "message"`
	Status  string `json: "message"`
}

func V1(app *application.App) *gin.Engine {
	router := gin.Default()
	userController := UserControllerV1{app: app}

	router.GET("/", func(c *gin.Context) {
		log.Debug().Msg(c.Request.Host)
		c.JSON(200, struct {
			Users string
		}{
			Users: "/users/",
		})
	})
	{
		group := router.Group("/users")
		group.GET("/", userController.ListUsers)
		group.POST("/", userController.CreateUser)

		group.GET("/:id/", userController.DetailUser)
		group.PUT("/:id/", userController.UpdateUser)
		group.DELETE("/:id/")
		group.PATCH("/:id/", userController.UpdateUser)
	}


	return router
}
