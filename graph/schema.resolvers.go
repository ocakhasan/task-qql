package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"io/ioutil"

	"github.com/pvormste/task-gql-full-stack/graph/generated"
	"github.com/pvormste/task-gql-full-stack/graph/model"
	"github.com/pvormste/task-gql-full-stack/schemaparser"
)

func (r *queryResolver) Heroes(_ context.Context) ([]model.Character, error) {
	// Return a hardcoded array of information here.
	heroes := []model.Character{
		model.Human{
			Name:          "Luke Skywalker",
			HasLightsaber: true,
		},
		model.Human{
			Name:          "Han Solo",
			HasLightsaber: false,
		},
		model.Droid{
			Name:            "C-3PO",
			PrimaryFunction: "Translator",
		},
	}

	return heroes, nil
}

func (r *queryResolver) Types(ctx context.Context) ([]*string, error) {
	schema, err := ioutil.ReadFile("./graph/schema.graphqls")
	if err != nil {
		return nil, err
	}
	typeNames, err := schemaparser.ParseTypes(schema)
	if err != nil {
		return nil, err
	}
	typeNamesResult := make([]*string, 0)
	for _, tn := range typeNames {
		tn := tn
		typeNamesResult = append(typeNamesResult, &tn)
	}

	return typeNamesResult, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
