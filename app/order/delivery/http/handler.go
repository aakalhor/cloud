package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"net/http"
	"order-pet/domain"
	pba "order-pet/pba"
)

type handler struct {
	u domain.OrderUsecase
	a pba.AuthServiceClient
}

func New(rg *gin.RouterGroup, u domain.OrderUsecase, conn *grpc.ClientConn) {
	a := pba.NewAuthServiceClient(conn)
	handler := &handler{
		u: u,
		a: a,
	}

	rg.POST("/register/:userId", handler.register)
	rg.GET("/get/:userId", handler.getOrders)
}

// create order godoc
// @Summary create order
// @Schemes
// @Description create order
// @Tags order
// @Accept json
// @Produce json
// @Param json	body RegisterOrder true  "item"
// @Param Authorization	header string 			true "Bearer jwtToken"
// @Success 200 {object} string
// @Failure      500
// @Failure      400
// @Router /order/register [post]
func (h *handler) register(ctx *gin.Context) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	token := authHeader[len(BEARER_SCHEMA):]
	res, err := h.a.Validate(ctx, &pba.ValidateRequest{Token: token})
	if err != nil {
		return
	}
	if res.Status != http.StatusOK {
		ctx.JSON(int(res.Status), gin.H{})
		return
	}
	username := ctx.Param("userId")
	var order RegisterOrder
	err = ctx.BindJSON(&order)

	if err != nil {
		return
	}
	productId, err := uuid.Parse(order.ProductId)

	_, err = h.u.RegisterOrder(ctx, username, productId, order.Date, order.Count)
	if err != nil {
		return
	}
	ctx.String(201, "")
}

// create order godoc
// @Summary get order
// @Schemes
// @Description get order
// @Tags order
// @Accept json
// @Produce json
// @Param userId	path string true  "User ID"
// @Param Authorization	header string 			true "Bearer jwtToken"
// @Success 200 {array} domain.Order
// @Failure      500
// @Failure      400
// @Router /order/get/:userId [get]
func (h *handler) getOrders(ctx *gin.Context) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	token := authHeader[len(BEARER_SCHEMA):]
	res, err := h.a.Validate(ctx, &pba.ValidateRequest{Token: token})
	if err != nil {
		return
	}
	if res.Status != http.StatusOK {
		ctx.JSON(int(res.Status), gin.H{})
		return
	}

	username := ctx.Param("username")
	orders, err := h.u.GetOrdersByUsername(ctx, username)
	if err != nil {
		return
	}
	//jsonOrders, err := json.Marshal(orders)
	//if err != nil {
	//	ctx.JSON(400, gin.H{})
	//	return
	//}
	ctx.JSON(200, &orders)
}
