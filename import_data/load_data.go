package main

import (
	"context"
	"encoding/xml"
	"fmt"
	f "herd/functions"
	h "herd/models"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	var days int64
	var herd h.Herd_Xml
	var oldInAges float64

	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "NumberOfDays", "fileName")
		return
	}
	xmlFile, err := os.Open(os.Args[2])
	days, _ = strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(byteValue, &herd)
	var listOfYak []interface{}

	for i := 0; i < len(herd.Labyak); i++ {
		oldInAges, _ = strconv.ParseFloat(herd.Labyak[i].Age, 64)
		
		var herds []h.Herd
		var age,_ = strconv.ParseFloat(herd.Labyak[i].Age,64)
		var lastShaved = float64(f.GetLastShaved(int(days), oldInAges))

		var herd = h.Herd{herd.Labyak[i].Name, age, lastShaved }
		herds = append(herds, herd)
		for _, herd := range herds {
			listOfYak = append(listOfYak, herd)
		}

	}
	ctx := context.Background()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	collection := client.Database("YakStore").Collection("herdCollection")
	deleteManyResult, err := collection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	insertManyResult, err := collection.InsertMany(ctx, listOfYak)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deleted Yak data before: ", deleteManyResult.DeletedCount)
	log.Println("Inserted Yak data: ", len(insertManyResult.InsertedIDs))
	
}
