package main

import (
	"context"
	"fmt"
	"strconv"

	pb "github.com/ronexlemon/godocker/common/api"
	"google.golang.org/grpc"
)

type Server  struct{
	pb.UnimplementedStoreServiceServer
}

func NewGRPCServiceHandler(grpcServer *grpc.Server){
	handler := &Server{}
	pb.RegisterStoreServiceServer(grpcServer,handler)
}

func (s *Server) GetMenu(menurequest *pb.MenuRequest,gr  grpc.ServerStreamingServer[pb.OrderMenu]) error {
	
	orders := []*pb.Order{
		{Id: "Order 1", Name: "Black Coffee"},
		{Id: "Order 2", Name: "Capuchino"},
		{Id: "Order 3", Name: "Black Tea"},


	}
	
	for index, order := range orders {
		orderMenu:= pb.OrderMenu{
			Id: fmt.Sprintf("%s",strconv.Itoa(index)),
			Orders: []*pb.Order{order},
		}
		if err := gr.Send(&orderMenu); err != nil {
			return err
			}
			}
			
			

	return nil
}
func (s *Server) PlaceOrder(ctx context.Context, order *pb.Order) (*pb.Receipt, error) {
	fmt.Println("Receive Order",order)
	return &pb.Receipt{
		Id: "Service Store Id 1234",
		Name: order.Name,
	},nil
}
func (s *Server) CheckStatus(ctx context.Context, rceipt *pb.Receipt) (*pb.OrderStatus, error) {
	fmt.Println("Check Status",rceipt)
	return &pb.OrderStatus{
		Id: "Order of the day 1",
		Status: "Order is ready",
	},nil
}
func (s *Server) CancelOrder(ctx context.Context, receipt *pb.Receipt) (*pb.Receipt, error) {
	fmt.Println("Cancel Order",receipt)
	return &pb.Receipt{
		Id: "Service Store Id 1234",
		Name: receipt.Name,
	},nil
}