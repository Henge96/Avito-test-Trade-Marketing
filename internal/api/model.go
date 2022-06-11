package api

type Event struct {
	Date   Date    `json:"date"`
	Views  uint    `json:"views"`
	Clicks uint    `json:"clicks"`
	Cost   float64 `json:"cost"`
}

type Date struct {
	DateStartedAt string `json:"date_started_at"`
	StartPeriod   string `json:"start_period"`
	EndPeriod     string `json:"end_period"`
}

type Request struct {
	Id     int     `json:"id"`
	Date   string  `json:"date_request"`
	Views  uint    `json:"views_request"`
	Clicks uint    `json:"clicks_request"`
	Cost   float64 `json:"cost_request"`
	Cpc    float64 `json:"cpc"`
	Cpm    float64 `json:"cpm"`
}
