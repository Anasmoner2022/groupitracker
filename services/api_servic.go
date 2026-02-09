package services

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"io"
	"net/http"
)

const BaseURL = "https://groupietrackers.herokuapp.com/api"

func FetchArtist() ([]models.Artist, error) {
	//* 1. Build a full url

	artistUrl := BaseURL + "/artists"
	//* 2. Make HTTP GET Request

	resp, err := http.Get(artistUrl)
	if err != nil {
		//! Network Error
		return nil, fmt.Errorf("Field to fetch Artist %w", err)
	}

	//* 3. Close Response body when done

	defer resp.Body.Close()

	//* 4. Check if the request was successful

	if resp.StatusCode != http.StatusOK {
		//! Server Returned an error
		return nil, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	//* 5. Read the response body (JSON Data)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		//! Couldn't read the response
		return nil, fmt.Errorf("Field to read response body: %w", err)
	}

	//* 6. Parse the JSON into our Go Struct

	var artist []models.Artist
	err = json.Unmarshal(body, &artist)
	if err != nil {
		//! JSON wan invalid or didn't match our struct
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return artist, nil

}
