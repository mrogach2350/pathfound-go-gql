package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/mrogach2350/pathfound_go/graph/generated"
	"github.com/mrogach2350/pathfound_go/graph/model"
	"github.com/mrogach2350/pathfound_go/helpers"
	"github.com/mrogach2350/pathfound_go/models"
)

func (r *mutationResolver) CreateGroup(ctx context.Context, input *model.NewGroup) (*models.Group, error) {
	return &models.Group{Name: *input.Name}, nil
}

func (r *queryResolver) Weapons(ctx context.Context) ([]*models.Weapon, error) {
	weapons := helpers.GetWeapons()
	return weapons, nil
}

func (r *queryResolver) Weapon(ctx context.Context, input *int) (*models.Weapon, error) {
	weapons := helpers.GetWeapons()
	idx := *input - 1
	if idx >= 0 && idx < len(weapons) {
		return weapons[idx], nil
	}
	return nil, errors.New("index is out of bounds")
}

func (r *queryResolver) Armors(ctx context.Context) ([]*models.Armor, error) {
	armor := helpers.GetArmor()
	return armor, nil
}

func (r *queryResolver) Armor(ctx context.Context, input *int) (*models.Armor, error) {
	armor := helpers.GetArmor()
	idx := *input - 1
	if idx >= 0 && idx < len(armor) {
		return armor[idx], nil
	}
	return nil, errors.New("index is out of bounds")
}

func (r *queryResolver) Groups(ctx context.Context) ([]*models.Group, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
