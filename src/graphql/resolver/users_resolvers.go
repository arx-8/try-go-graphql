package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/arx-8/try-go-graphql/src/graphql/gql_model"
	"github.com/arx-8/try-go-graphql/src/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input gql_model.NewUser) (*gql_model.User, error) {
	record := model.User{
		Name: input.Name,
	}
	if err := r.DB.Create(&record).Error; err != nil {
		return nil, err
	}

	res := record.ToGqlModel()
	return &res, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*gql_model.User, error) {
	idn, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	var u model.User
	if err := r.DB.Find(&u, idn).Error; err != nil {
		return nil, err
	}

	res := u.ToGqlModel()
	return &res, nil
}
