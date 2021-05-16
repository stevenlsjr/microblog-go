package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"microblog/application"
	"microblog/models"
)

type QueryListUsers struct {
	Limit  int64 `form:"limit" json:"limit" xml:"limit" default:"10"`
	Offset int64 `form:"offset" json:"offset" xml:"offset" default:"0"`
}

type CreateUserV1 struct {
	ID       *string `json:"id" form:"id" xml:"id"`
	Username string  `json:"username" form:"username" xml:"username"`
}

type UpdateUserV1 struct {
	ID       string `json:"id" form:"id" xml:"id"`
	Username string `json:"username" form:"username" xml:"username"`
}

type PartialUpdateUserV1 struct {
	ID       *string `json:"id" form:"id" xml:"id"`
	Username *string `json:"username" form:"username" xml:"username"`
}

type UserControllerV1 struct {
	app *application.App
}

type listUsersV1 struct {
	PaginationParams
	Results []models.BlogUser `json:"results"`
}

func (u *UserControllerV1) ListUsers(context *gin.Context) {
	query := QueryListUsers{}
	err := context.BindQuery(&query)
	if err != nil {
		_ = context.AbortWithError(400, err)
		return
	}
	var users []models.BlogUser
	var count int64
	err = u.app.Db.Transaction(func(tx *gorm.DB) error {
		tx.Limit(int(query.Limit)).Offset(int(query.Offset)).Find(&users)
		tx.Model(&models.BlogUser{}).Count(&count)
		return nil
	})
	if err != nil {
		_ = context.AbortWithError(500, err)
		return
	}

	body := listUsersV1{
		PaginationParams: PaginationParams{
			Limit:  query.Limit,
			Offset: query.Offset,
			Count:  count,
		},
		Results: users,
	}
	context.JSON(200, body)
}

func (u *UserControllerV1) DetailUser(context *gin.Context) {
	id := context.Param("id")
	user, err := u.app.Users.FindById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		_ = context.AbortWithError(404, err)
		log.Error().Err(err)
		return
	}
	context.JSON(200, user)
}

func (u *UserControllerV1) CreateUser(context *gin.Context) {
	var data CreateUserV1
	err := context.BindJSON(&data)
	if err != nil {
		_ = context.AbortWithError(400, err)
		return
	}
	if data.ID == nil {
		defaultId := uuid.NewString()
		data.ID = &defaultId
	}
	user := &models.BlogUser{
		UuidModel: *models.NewUuidModel(*data.ID),
		Username:  data.Username,
	}
	user, err = u.app.Users.Create(user)
	if err != nil {
		var duplicateKey *application.ErrDuplicateKey
		if errors.As(err, &duplicateKey) {
			_ = context.Error(duplicateKey)
			context.AbortWithStatusJSON(400, ResponseError{Status: ResponseBadRequest,
				Message: fmt.Sprintf("Duplicate key: %s", duplicateKey.ColumnOrConstraint)})
		} else {
			_ = context.Error(err)
			context.AbortWithStatusJSON(500, ResponseError{Status: ResponseServerError})
		}
		return
	} else {
		context.JSON(201, user)
	}
}

func (u *UserControllerV1) UpdateUser(context *gin.Context) {

	var updateUser UpdateUserV1
	if err := context.Bind(&updateUser); err != nil {
		return
	}
	id := context.Param("id")

	if updateUser.ID != id {
		context.AbortWithStatusJSON(400, ResponseError{
			Message: fmt.Sprintf("path id '%s' does not match id in body: '%s'", id, updateUser.ID),
			Status:  ResponseBadRequest,
			Data:    "",
		})
	}


}
