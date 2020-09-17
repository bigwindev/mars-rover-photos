package cache

import (
	"encoding/json"

	"mars-rover-photos/models"
)

func SetMarsPhotos(date string, camera string, photos string) error {
	// Get cache connection from pool
	conn, err := GetConn()
	if err != nil {
		return err
	}

	// Store in cache
	err = Set(conn, date+":"+camera, photos)

	// Close redis connection
	conn.Close()

	if err != nil {
		return err
	}

	return nil
}

func GetMarsPhotos(date string, camera string) (*models.Photos, error) {
	// Get cache connection from pool
	conn, err := GetConn()
	if err != nil {
		return nil, err
	}

	// Get from cache
	s, err := Get(conn, date+":"+camera)

	// Close redis connection
	conn.Close()

	if err != nil {
		return nil, err
	}

	// Unmarshal response from cache
	photos := &models.Photos{}
	err = json.Unmarshal([]byte(s), photos)
	if err != nil {
		return nil, err
	}

	return photos, nil
}
