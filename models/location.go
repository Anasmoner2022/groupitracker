package models

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     *Date    `json:"dates"`
}
