package models

type Herd struct {
	Name *string  `json:"name"`
	Age  *float64 `json:"age"`
	Sex  *string  `json:"sex"`
}
