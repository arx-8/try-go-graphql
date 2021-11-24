package model

import (
	"fmt"

	"github.com/arx-8/try-go-graphql/src/graphql/gql_model"
)

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (e User) ToGqlModel() gql_model.User {
	return gql_model.User{
		ID:   fmt.Sprintf("%d", e.ID),
		Name: e.Name,
	}
}
