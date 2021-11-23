package gorm_model

import (
	"fmt"

	"github.com/arx-8/try-go-graphql/src/entity"
	"github.com/arx-8/try-go-graphql/src/graphql/model"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUserFromEntity(e *entity.User) *model.User {
	return &model.User{
		ID:   fmt.Sprintf("%d", e.ID),
		Name: e.Name,
	}
}
