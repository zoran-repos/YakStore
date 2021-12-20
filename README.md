# YakStore
1. For checking first task - follow next steps:
   -  go mod tidy
   -  go build task1.go
   -  ./task1 13 herd.xml
   -  < response is here >
   -  ./task1 14 herd.xml
   -  < response is here >
1.1 For data manipulation we use MongoDb as a document store - this functionality is only for presentation - not for production
    You can install MongoDb https://docs.mongodb.com/manual/installation/ or with docker 
    From mongo(shell) you can run those commands for creation db and collection
    - use YakStore
      db.createCollection("herdCollection")
