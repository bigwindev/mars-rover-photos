package models

// rover camera resource
type Camera struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	RoverId  int    `json:"rover_id"`
	FullName string `json:"full_name"`
}
