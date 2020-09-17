package models

// mars rover photo resource
type Photo struct {
	Id int `json:"id"`
	Sol int `json:"sol"`
	Camera Camera `json:"camera"`
	ImgSrc string `json:"img_src"`
	EarthDate string `json:"earth_date"`
	Rover Rover `json:"rover"`
}
