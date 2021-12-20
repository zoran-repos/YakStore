package main

import (
	"context"
	"log"

	handlers "herd/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var herdsHandler *handlers.HerdsHandler

func init() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	collection := client.Database("YakStore").Collection("herdCollection")
	herdsHandler = handlers.NewHerdsHandler(ctx, collection)
}

func main() {
	router := gin.Default()
	router.GET("/yak-shop/herd/:days", herdsHandler.ListHerdsHandler)
	router.GET("/yak-shop/stock/:days", herdsHandler.ListStockHandler)
	router.POST("/yak-shop/order/:days", herdsHandler.OrdersHandler)
	router.Run()
}
