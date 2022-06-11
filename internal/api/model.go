package api

type Event struct {
	Date   string `json:"date"`
	Views  uint   `json:"views"`
	Clicks uint   `json:"clicks"`
	Cost   string `json:"cost"`
}

type EventPeriod struct {
	StartedAt string `json:"started_at"`
	EndedAt   string `json:"ended_at"`
}

type Response struct {
	Date   string `json:"date"`
	Views  uint   `json:"views"`
	Clicks uint   `json:"clicks"`
	Cost   string `json:"cost"`
	Cpc    string `json:"cpc"`
	Cpm    string `json:"cpm"`
}
