package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"net/http"
	pba "shop-pet/pba"
	//"github.com/google/uuid"
	//"net/http"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"shop-pet/domain"
	//"time"
)

type Handler struct {
	u domain.ShopUsecase
	a pba.AuthServiceClient
}

// item godoc
// @Summary register item
// @Schemes
// @Description create item
// @Tags shop
// @Accept json
// @Produce json
// @Param json	body RegisterItem true  "item"
// @Param Authorization	header string 			true "Bearer jwtToken"
// @Success 200 {object} string
// @Failure      500
// @Failure      400
// @Router /shop/register-item [post]
func (h *Handler) registerItem(ctx *gin.Context) {
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
	var req RegisterItem
	err = ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "can not bind json body"})
	}
	item, err := h.u.RegisterItem(ctx, req.CategoryName, req.Name, req.Sex, req.Price, req.Count)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
	}
	ctx.JSON(http.StatusCreated, &item)

}

// category godoc
// @Summary register category
// @Schemes
// @Description create category
// @Tags shop
// @Accept json
// @Produce json
// @Param json	body RegisterCategory true  "item"
// @Param Authorization	header string 			true "Bearer jwtToken"
// @Success 200 {object} string
// @Failure      500
// @Failure      400
// @Router /shop/register-category [post]
func (h *Handler) registerCategory(ctx *gin.Context) {
	fmt.Println("here!")
	const BEARER_SCHEMA = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	token := authHeader[len(BEARER_SCHEMA):]
	fmt.Println(token)
	res, err := h.a.Validate(ctx, &pba.ValidateRequest{Token: token})
	fmt.Println(res, err)
	if err != nil {
		return
	}
	if res.Status != http.StatusOK {
		fmt.Println(res.Status, "<<<")
		ctx.JSON(int(res.Status), gin.H{})
		return
	}
	fmt.Println("here ...")
	var req RegisterCategory
	err = ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "can not bind json body"})
	}
	category, err := h.u.RegisterCategory(ctx, req.Name, req.Desc)
	ctx.JSON(http.StatusCreated, &category)
}

// add item godoc
// @Summary add item
// @Schemes
// @Description add item
// @Tags shop
// @Accept json
// @Produce json
// @Param json	body AddItem true  "item"
// @Param Authorization	header string 			true "Bearer jwtToken"
// @Success 200 {object} string
// @Failure      500
// @Failure      400
// @Router /shop/add-item [post]
func (h *Handler) addToItem(ctx *gin.Context) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	token := authHeader[len(BEARER_SCHEMA):]
	fmt.Println(token)
	res, err := h.a.Validate(ctx, &pba.ValidateRequest{Token: token})
	fmt.Println(res, err)
	if err != nil {
		return
	}
	if res.Status != http.StatusOK {
		fmt.Println(res.Status, "<<<")
		ctx.JSON(int(res.Status), gin.H{})
		return
	}
	var req AddItem
	err = ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "can not bind json body"})
	}
	id := uuid.MustParse(req.Id)
	err = h.u.AddToItem(ctx, id, req.Count)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": "can not add to this item"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

// item godoc
// @Summary register item
// @Schemes
// @Description create item
// @Tags shop
// @Accept json
// @Produce json
// @Param json	body RegisterCategory true  "item"
// @Param Authorization	header string 			true "Bearer jwtToken"
// @Success 200 {object} string
// @Failure      500
// @Failure      400
// @Router /shop/get-item [get]
func (h *Handler) getItem(ctx *gin.Context) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	token := authHeader[len(BEARER_SCHEMA):]
	fmt.Println(token)
	res, err := h.a.Validate(ctx, &pba.ValidateRequest{Token: token})
	fmt.Println(res, err)
	if err != nil {
		return
	}
	if res.Status != http.StatusOK {
		fmt.Println(res.Status, "<<<")
		ctx.JSON(int(res.Status), gin.H{})
		return
	}
	id := ctx.Param("itemId")
	itemId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "can not convert to id"})
		return
	}

	item, err := h.u.GetItem(ctx, itemId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "can no get item"})
		return
	}
	ctx.JSON(http.StatusOK, &item)
}

func New(rg *gin.RouterGroup, u domain.ShopUsecase, conn *grpc.ClientConn) (*Handler, error) {
	a := pba.NewAuthServiceClient(conn)
	handler := Handler{
		u: u,
		a: a,
	}
	rg.POST("/register-item", handler.registerItem)
	rg.POST("/register-category", handler.registerCategory)
	rg.POST("/add-item", handler.addToItem)
	rg.GET("/get-item/:itemId", handler.getItem)

	rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return &handler, nil
}
