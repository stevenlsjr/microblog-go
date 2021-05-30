package models

type User struct {
	UuidModel
	Username string    `gorm:"unique;index" json:"username" form:"username" xml:"username" faker:"username"`
	Email    string    `gorm:"unique" json:"email" form:"email" xml:"email" faker:"email"`
	Messages []Message `gorm:"foreignKey:AuthorId" json:"messages" default:"[]"`
}
