package common

import (
	"mars-rover-photos/models"
)

var (
	RoverCameras = []models.Camera{
		models.Camera{
			Name:     "FHAZ",
			FullName: "Front Hazard Avoidance Camera",
		},
		models.Camera{
			Name:     "RHAZ",
			FullName: "Rear Hazard Avoidance Camera",
		},
		models.Camera{
			Name:     "MAST",
			FullName: "Mast Camera",
		},
		models.Camera{
			Name:     "CHEMCAM",
			FullName: "Chemistry and Camera Complex",
		},
		models.Camera{
			Name:     "MAHLI",
			FullName: "Mars Hand Lens Imager",
		},
		models.Camera{
			Name:     "MARDI",
			FullName: "Mars Descent Imager",
		},
		models.Camera{
			Name:     "NAVCAM",
			FullName: "Navigation Camera",
		},
		models.Camera{
			Name:     "PANCAM",
			FullName: "Panoramic Camera",
		},
		models.Camera{
			Name:     "MINITES",
			FullName: "Miniature Thermal Emission Spectrometer (Mini-TES)",
		},
	}
)
