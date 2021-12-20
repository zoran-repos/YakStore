package models

type Herd_mongo struct {
	Name            string  `bson:"name"`
	Age             float64 `bson:"age"`
	Age_Last_Shaved float64 `bson:"age_last_shaved"`
}
