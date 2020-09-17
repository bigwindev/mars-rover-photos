package nasa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"mars-rover-photos/config"
	"mars-rover-photos/models"
	"mars-rover-photos/utils/cache"
)

// GetMarsPhotos parses the response from NASA Mars Rover Phtos API
func GetMarsPhotos(date string, camera string) (*models.Photos, error) {
	// Create URL data
	urlData := url.Values{}
	urlData.Set("earth_date", date)
	if camera != "" && camera != "all" {
		urlData.Set("camera", camera)
	}
	urlData.Set("api_key", config.Nasa.APIKey)

	url, _ := url.ParseRequestURI(config.Nasa.APIURL)
	url.Path = config.Nasa.MRPPath
	url.RawQuery = urlData.Encode()
	encodedUrl := fmt.Sprintf("%v", url)

	fmt.Printf("URL: %v\n", encodedUrl)

	// Send a request to NASA API
	client := &http.Client{}
	req, err := http.NewRequest("GET", encodedUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(urlData.Encode())))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse response from API
	photos := &models.Photos{}
	err = json.Unmarshal(jsonData, photos)
	if err != nil {
		return nil, err
	}

	// Store in cache
	err = cache.SetMarsPhotos(date, camera, string(jsonData))
	if err != nil {
		log.Printf("mars-rover-photos: failed to store in cache: %v\n", err)
	}

	return photos, nil
}
