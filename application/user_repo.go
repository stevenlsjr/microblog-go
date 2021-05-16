package application

import (
	"errors"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
	"microblog/models"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) FindById(id string) (*models.BlogUser, error){
	user := models.BlogUser{}
	query := u.db.Model(&models.BlogUser{}).
		Find(&user).
		Where("ID", id)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return &user, nil
	}
}

func (u *UserRepo) Create(user *models.BlogUser) (*models.BlogUser, error) {
	query := u.db.Create(user)
	var pgErr *pgconn.PgError
	if errors.As(query.Error, &pgErr) {
		return nil, ToAppError(pgErr)
	} else if err := query.Error; err != nil {
		return nil, err
	}
	return user, nil
}

type UpdateUser struct {
	Username *string
}

func (u *UserRepo) Update(id string, fields UpdateUser) (*models.BlogUser, error){
	return nil, errors.New("not implemented")
}