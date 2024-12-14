package model

type Profile struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Lat      string `json:"lat"`
	Long     string `json:"long"`
	DietType string `json:"diet_type"`
	Age      int    `json:"age"`
}
