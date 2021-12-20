package handlers

import (
	"net/http"
	"strconv"

	f "herd/functions"
	h "herd/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type HerdsHandler struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewHerdsHandler(ctx context.Context, collection *mongo.Collection) *HerdsHandler {
	return &HerdsHandler{
		collection: collection,
		ctx:        ctx,
	}
}

func (handler *HerdsHandler) ListHerdsHandler(c *gin.Context) {
	cur, err := handler.collection.Find(handler.ctx, bson.M{})
	id, _ := strconv.Atoi(c.Param("days"))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(handler.ctx)

	herds := make([]h.Herd_mongo, 0)
	for cur.Next(handler.ctx) {
		var herd h.Herd_mongo
		cur.Decode(&herd)
		herd.Age_Last_Shaved = float64(f.GetLastShaved(id, herd.Age))
		herds = append(herds, herd)
	}
	c.JSON(http.StatusOK, herds)
}

func (handler *HerdsHandler) ListStockHandler(c *gin.Context) {

	var tillDateStock, milkStock float64
	var tillDateSkinStock, skinStock int

	cur, err := handler.collection.Find(handler.ctx, bson.M{})
	id, _ := strconv.Atoi(c.Param("days"))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(handler.ctx)

	for cur.Next(handler.ctx) {
		var herd h.Herd_mongo

		cur.Decode(&herd)
		tillDateStock = f.GetTotalMilk(id, herd.Age)
		milkStock += tillDateStock

		tillDateSkinStock = f.GetSkin(id, herd.Age)
		skinStock += tillDateSkinStock
	}

	var stocks = []h.Stock{
		{Milk: milkStock, Skins: skinStock},
	}

	c.JSON(http.StatusOK, stocks)
}

func (handler *HerdsHandler) OrdersHandler(c *gin.Context) {

	var tillDateStock, milkStock float64
	var tillDateSkinStock, skinStock int

	id, _ := strconv.Atoi(c.Param("days"))
	var requestOrder h.RequestOrder
	if err := c.ShouldBindJSON(&requestOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cur, err := handler.collection.Find(handler.ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(handler.ctx)

	for cur.Next(handler.ctx) {
		var herd h.Herd_mongo

		cur.Decode(&herd)
		tillDateStock = f.GetTotalMilk(id, herd.Age)
		milkStock += tillDateStock

		tillDateSkinStock = f.GetSkin(id, herd.Age)
		skinStock += tillDateSkinStock
	}

	if (requestOrder.Order.Milk >= milkStock) && (requestOrder.Order.Skins >= skinStock) {
		var order_result = h.Stock{}
		c.JSON(http.StatusNotFound, order_result)
	}
	if (requestOrder.Order.Milk >= milkStock) && (requestOrder.Order.Skins < skinStock) {
		var order_result = h.Stock{Skins: requestOrder.Order.Skins}
		c.JSON(http.StatusPartialContent, order_result)
	}
	if (requestOrder.Order.Milk < milkStock) && (requestOrder.Order.Skins < skinStock) {
		var order_result = h.Stock{Milk: requestOrder.Order.Milk, Skins: requestOrder.Order.Skins}
		c.JSON(http.StatusOK, order_result)
	}
	if (requestOrder.Order.Milk < milkStock) && (requestOrder.Order.Skins >= skinStock) {
		var order_result = h.Stock{Milk: requestOrder.Order.Milk}
		c.JSON(http.StatusPartialContent, order_result)
	}

}
