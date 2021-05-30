package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"microblog/application"
	"microblog/models"
)

type UserControllerV1 struct {
	app *application.App
}

type userUriParams struct {
	ID string `uri:"id" binding:"required"`
}

func NewUsersV1(app *application.App) *UserControllerV1 {
	return &UserControllerV1{app: app}
}

func (u *UserControllerV1) db() *gorm.DB {
	return u.app.Db
}

type UserList struct {
	PaginationParams
	Objects []models.User `json:"objects"`
}

func (u *UserControllerV1) List(ctx *gin.Context) {
	var queryParams LimitOffset
	if err := ctx.BindQuery(&queryParams); err != nil {
		return
	}
	var resp UserList

	q := u.db().Model(&models.User{}).
		Order("id").
		Limit(int(queryParams.Limit)).
		Offset(int(queryParams.Offset))
	if q.Find(&resp.Objects).Error != nil {
		_ = ctx.AbortWithError(500, q.Error)
		return
	}
	if q.Count(&resp.Count).Error != nil {
		_ = ctx.AbortWithError(500, q.Error)
		return
	}

	ctx.JSON(200, resp)
}

func (u *UserControllerV1) Detail(ctx *gin.Context) {
	var params userUriParams
	var user models.User
	if err := ctx.BindUri(&params); err != nil {
		return
	}

	if res := u.db().Take(&user, "id = ?", params.ID); res.Error != nil {
		var status int
		switch {
		case errors.Is(res.Error, gorm.ErrRecordNotFound):
			status = 404

		default:
			status = 500
		}
		_ = ctx.AbortWithError(status, res.Error)
		return
	}

	ctx.JSON(200, user)

}

func (u *UserControllerV1) Create(ctx *gin.Context) {

}

func (u *UserControllerV1) Update(ctx *gin.Context) {

}

func (u *UserControllerV1) Delete(ctx *gin.Context) {

}

func (u *UserControllerV1) PartialUpdate(ctx *gin.Context) {

}
