package user

import "go-blog/app/models"

type User struct {
	models.BaseModel

	Name string `json:"name,omitempty"`
	Email string `json:"-"`
	Phone string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
