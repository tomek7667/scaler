package domain

type Entry struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}

type Scale struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	ScalePassword string  `json:"scalePassword"`
	Entries       []Entry `json:"entries"`
}
