package models

type Herd struct {
	Name            string  `json:"name"`
	Age             float64 `json:"age"`
	Age_Last_Shaved float64 `json:"age-last-shaved"`
}
