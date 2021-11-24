package model

import (
	"fmt"

	"github.com/arx-8/try-go-graphql/src/graphql/gql_model"
)

type Todo struct {
	ID     uint   `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID uint   `json:"user"`
}

func (e Todo) ToGqlModel() gql_model.Todo {
	return gql_model.Todo{
		ID:     fmt.Sprintf("%d", e.ID),
		Text:   e.Text,
		Done:   e.Done,
		UserID: fmt.Sprintf("%d", e.UserID),
	}
}
