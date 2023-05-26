package user

import (
	"go-blog/app/models"
	"go-blog/pkg/database"
)

type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

func (userModel *User) Create() {
	database.DB.Create(&userModel)
}
