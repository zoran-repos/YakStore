# YakStore
1. For checking first task - follow next steps:
    -  go mod tidy
    -  go build task1.go
    -  ./task1 13 herd.xml
    -  < response is here >
    -  ./task1 14 herd.xml
    -  < response is here >
   
1.1 For data manipulation we use MongoDb as a document store - this functionality is only for presentation - not for production.

    You can install MongoDb https://docs.mongodb.com/manual/installation/ or with docker 
    
    From mongo(shell) you can run those commands for creation db and collection
    -  use YakStore 
    -  db.createCollection("herdCollection")
   
    In folder import_data we have a tool for loading data - load_data.go
    We upload data with syntax
    -  ./load_data 13 ../herd.xml

2. For checking how 2nd task is resolved - follow next steps
   -  go mod tidy
   In folder endpoints 
      go run main.go
   You can test with postman or any other tools: 
   - 
   - http://localhost:8080/yak-shop/stock/13
   - http://localhost:8080/yak-shop/stock/14
   - http://localhost:8080/yak-shop/herd/13
   - http://localhost:8080/yak-shop/herd/14
     < - here we have diferent response for age-last-shaved for Betty-1 >

