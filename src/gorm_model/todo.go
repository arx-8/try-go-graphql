package gorm_model

import (
	"fmt"

	"github.com/arx-8/try-go-graphql/src/entity"
	"github.com/arx-8/try-go-graphql/src/graphql/model"
)

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"user"`
}

func NewTodoFromEntity(e *entity.Todo) *model.Todo {
	return &model.Todo{
		ID:     fmt.Sprintf("%d", e.ID),
		Text:   e.Text,
		Done:   e.Done,
		UserID: fmt.Sprintf("%d", e.UserID),
	}
}
