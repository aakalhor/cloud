package grpc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"shop-pet/domain"
	pb "shop-pet/pb/grpc"
	"time"
)

type handler struct {
	pb.UnimplementedShopServiceServer
	u domain.ShopUsecase
}

func New(s grpc.ServiceRegistrar, u domain.ShopUsecase) {
	h := &handler{
		u: u,
	}
	pb.RegisterShopServiceServer(s, h)
}

func (h *handler) Create(ctx context.Context, req *pb.CreateOrderReq) (*pb.CreateOrderRes, error) {
	fmt.Println("in creaaaete")
	fmt.Println(req)
	id, _ := uuid.Parse(req.Id)
	username := req.Username
	pId, _ := uuid.Parse(req.ProductId)
	date := time.Now()
	count := int(req.Count)

	id, err := h.u.RegisterOrder(ctx, id, username, pId, date, count)
	fmt.Println("id", id, err)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderRes{
		Status: 1,
		Error:  "",
	}, nil
}

func (h *handler) Cancel(ctx context.Context, req *pb.CancelOrderReq) (*pb.CancelOrderRes, error) {
	id, _ := uuid.Parse(req.OrderId)
	h.u.CancelOrder(ctx, id)
	return nil, nil
}

func (h *handler) GetByUsername(ctx context.Context, req *pb.GetByUsernameReq) (*pb.GetByUsernameRes, error) {
	return nil, nil
}
