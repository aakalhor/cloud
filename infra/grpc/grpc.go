package grpc

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"order-pet/app/order/adapter"
	"order-pet/domain"
	pb "order-pet/pb/grpc"
	"time"
)

type handler struct {
	sc pb.ShopServiceClient
}

func New(client *grpc.ClientConn) adapter.Adapter {
	sc := pb.NewShopServiceClient(client)
	return &handler{
		sc: sc,
	}
}

func (h *handler) CreateOrder(ctx context.Context, username string, orderId uuid.UUID, productId uuid.UUID, count int, date time.Time) error {
	pId := productId.String()
	oId := orderId.String()

	_, err := h.sc.Create(ctx, &pb.CreateOrderReq{
		ProductId: pId,
		Username:  username,
		Date:      date.String(),
		Count:     int64(count),
		Id:        oId,
	})
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) CancelOrder(ctx context.Context, username string, orderId uuid.UUID) (returnAmount int, err error) {
	oId := orderId.String()

	res, err := h.sc.Cancel(ctx, &pb.CancelOrderReq{
		OrderId:  oId,
		Username: username,
	})
	remainAmount := res.RemainAmount
	if err != nil {
		return 0, err
	}
	return int(remainAmount), nil
}

func (h *handler) GetOrders(ctx context.Context, username string) (*[]domain.Order, error) {

	_, err := h.sc.GetByUsername(ctx, &pb.GetByUsernameReq{
		Username: username,
	})

	//TODO:fix it
	return nil, err
	//if err != nil {
	//	return 0, err
	//}
	//return int(remainAmount), nil
}
