package models

type RequestOrder struct {
	Customer string `json:"customer"`
	Order    struct {
		Milk  float64 `json:"milk"`
		Skins int `json:"skins"`
	} `json:"order"`
}