package cmd

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"order-pet/app/order/delivery/http"
	"order-pet/app/order/usecase"
	_ "order-pet/docs"
	"order-pet/domain"
	order_adapter "order-pet/infra/grpc"
)

var orderUsecase domain.OrderUsecase

func boot(gp *grpc.ClientConn) {

	orderAdapter := order_adapter.New(gp)

	orderUsecase = usecase.New(orderAdapter)

}

func Boot() {
	//orderAdapter, err := order_adapter.New(client, clientTrigger, kp)
	conn, err := grpc.Dial("localhost:10000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	boot(conn)
	// *****run http server*****

	router := gin.Default()
	authConn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	http.New(router.Group("order/"), orderUsecase, authConn)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8085")

	//_, err = t.GRPCSevrer("order-matching", func(server *grpc.Server) {
	//	_, err := order_grpc.New(server, orderUsecase)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//})

}
