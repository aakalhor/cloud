package cmd

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	shop_grpc "shop-pet/app/shop/delivery/grpc"
	"shop-pet/app/shop/delivery/http"
	"shop-pet/app/shop/repository"
	"shop-pet/app/shop/usecase"
	_ "shop-pet/docs"
	"shop-pet/domain"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(repository.Order{}, repository.Item{}, repository.Category{})

}

var shopUsecase domain.ShopUsecase

func boot(db *gorm.DB) {
	var err error
	shopRepository := repository.New(db)
	if err != nil {
		panic(err)
	}
	//shopAdapter, err := shop_adapter.New(kp)

	shopUsecase = usecase.New(shopRepository)

}

func Boot() {
	dsn := "host=localhost user=shop password=password dbname=shop port=5000 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//db := s.Database["main"]
	migrate(db)

	//orderAdapter, err := order_adapter.New(client, clientTrigger, kp)

	boot(db)
	//authMiddleware, err := AuthMiddleware()
	authMiddleware, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	// *****run http server*****
	router := gin.Default()

	_, err = http.New(router.Group("shop/"), shopUsecase, authMiddleware)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err != nil {
		panic(err)
	}

	go router.Run(":8081")
	lis, err := net.Listen("tcp", ":10000")
	grpcServer := grpc.NewServer()

	shop_grpc.New(grpcServer, shopUsecase)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
	//_, err = t.GRPCSevrer("order-matching", func(server *grpc.Server) {
	//	_, err := order_grpc.New(server, orderUsecase)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//})

}
