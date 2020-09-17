package controllers

import (
	"log"
  "fmt"
	"time"

	"mars-rover-photos/utils/nasa"
	"mars-rover-photos/utils/cache"
	"mars-rover-photos/common"
	"mars-rover-photos/models"

	"github.com/kataras/iris/v12"
)

// GetPhotos returns photos of provided date
func GetPhotos(ctx iris.Context) {
	// Get and parse earth date parameter to "YYYY-MM-DD" format
	dateParam := ctx.FormValue("earth_date")
	var date time.Time
	var dateStr string
	if dateParam == "" {
		date = time.Now()
		dateStr = date.Format("2006-01-02")
	} else {
		date, err := time.Parse("2006-01-02", dateParam)
		if err != nil {
			log.Printf("mars-rover-photos: time coverting error: %v\n", err)
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		}
		dateStr = date.Format("2006-01-02")
	}

	// Get camera parameter
	camera := ctx.FormValue("camera")

	// Get delta days
	days := time.Now().Sub(date).Hours() / 24

	var res *models.Photos
	var err error
	if days > 7 { // If more than a week
		// Get mars rover photos from cash
		res, err = cache.GetMarsPhotos(dateStr, camera)

		if err != nil {
			log.Printf("mars-rover-photos: failed to get data from cache with error: %v\n", err)

			// Get mars rover photos from NASA
			res, err = nasa.GetMarsPhotos(dateStr, camera)
		}
	} else {
		// Get mars rover photos from NASA
		res, err = nasa.GetMarsPhotos(dateStr, camera)
	}

	ctx.View("header.html")

	if err != nil {
		log.Printf("mars-rover-photos: failed to get nasa API response with error: %v\n", err)
		ctx.ViewData("Message", fmt.Sprintf("%v", err))
    ctx.View("error.html")
    return
	}

	ctx.ViewData("Date", dateStr)
	ctx.ViewData("Cameras", common.RoverCameras)
	ctx.ViewData("Camera", camera)
  ctx.View("form.html")

	// If no photos, then show empty message
	if len(res.Photos) == 0 {
    ctx.View("empty.html")
    return
	}

  ctx.ViewData("Photos", res.Photos)
  ctx.View("index.html")
}
