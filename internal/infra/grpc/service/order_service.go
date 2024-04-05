package service

import (
	"context"

	"github.com/beriloqueiroz/desafio-clean-arch/internal/infra/grpc/pb"
	"github.com/beriloqueiroz/desafio-clean-arch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase usecase.ListOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrderUseCase usecase.ListOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrderUseCase: listOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) (ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*ListOrderResponse, error){
	dto := usecase.ListOrderInputDTO{
		Page: in.Page,
		PageSize: in.PageSize,
	}
	output, err := s.ListOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	var listOrders []pb.ItemListOrderResponse
	
	for _,out := range output {
		listOrders = append(listOrders, pb.ItemListOrderResponse{
			Id:         out.ID,
			Price:      float32(out.Price),
			Tax:        float32(out.Tax),
			FinalPrice: float32(out.FinalPrice),
		})
	} 

	return &pb.ListOrderResponse{
		ItemListOrderResponse: listOrders,
	}, nil
}