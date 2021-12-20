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
