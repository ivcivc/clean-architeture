package service

import (
	"context"

	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/pb"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
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

func (s *OrderService) ListOrders(ctx context.Context, in *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	listOrdersUseCase := usecase.NewListOrdersUseCase(s.CreateOrderUseCase.OrderRepository)
	orders, err := listOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var responseOrders []*pb.Order
	for _, order := range orders {
		responseOrders = append(responseOrders, &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}
	return &pb.ListOrdersResponse{Orders: responseOrders}, nil
}
