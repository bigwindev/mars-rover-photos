package models

// rover resource
type Rover struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	LandingDate string `json:"landing_date"`
	LaunchDate  string `json:"launch_date"`
	Status      string `json:"status"`
}
