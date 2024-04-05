package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/beriloqueiroz/desafio-clean-arch/internal/infra/graph/model"
	"github.com/beriloqueiroz/desafio-clean-arch/internal/usecase"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.OrderInput) (*model.Order, error) {
	dto := usecase.OrderInputDTO{
		ID:    input.ID,
		Price: float64(input.Price),
		Tax:   float64(input.Tax),
	}
	output, err := r.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &model.Order{
		ID:         output.ID,
		Price:      float64(output.Price),
		Tax:        float64(output.Tax),
		FinalPrice: float64(output.FinalPrice),
	}, nil
}

// ListOrders is the resolver for the listOrders field.
func (r *mutationResolver) ListOrders(ctx context.Context, in *model.ListInput) ([]*model.Order, error) {
	dto := usecase.ListOrderInputDTO{
		Page:     in.Page,
		PageSize: in.PageSize,
	}
	output, err := r.ListOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	var listOrders []*model.Order

	for _, out := range output {
		listOrders = append(listOrders, &model.Order{
			ID:         out.ID,
			Price:      float64(out.Price),
			Tax:        float64(out.Tax),
			FinalPrice: float64(out.FinalPrice),
		})
	}

	return listOrders, nil

}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
