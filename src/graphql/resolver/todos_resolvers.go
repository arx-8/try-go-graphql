package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/arx-8/try-go-graphql/src/graphql/gql_model"
	"github.com/arx-8/try-go-graphql/src/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input gql_model.NewTodo) (*gql_model.Todo, error) {
	userID, err := strconv.Atoi(input.UserID)
	if err != nil {
		return nil, err
	}
	record := model.Todo{
		Text:   input.Text,
		UserID: uint(userID),
	}
	if err := r.DB.Create(&record).Error; err != nil {
		return nil, err
	}

	res := record.ToGqlModel()
	return &res, nil
}

func (r *queryResolver) Todos(ctx context.Context, userID string) ([]*gql_model.Todo, error) {
	userID_, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}
	var records []model.Todo
	if err := r.DB.Where("user_id", userID_).Find(&records).Error; err != nil {
		return nil, err
	}

	todos := []*gql_model.Todo{}
	for _, record := range records {
		g := record.ToGqlModel()
		todos = append(todos, &g)
	}

	return todos, nil
}
