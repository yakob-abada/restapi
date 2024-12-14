package model

type Explore struct {
	ID             int    `json:"id"`
	DietType       string `json:"diet_type"`
	AgeFrom        int    `json:"age_from"`
	AgeTo          int    `json:"age_to"`
	DistanceRadius int    `json:"distance_radius"`
	Gender         string `json:"gender"`
}
