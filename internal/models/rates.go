package models

type Rates struct {
	Ask       float64 `json:"ask"`
	Bid       float64 `json:"bid"`
	Timestamp string  `json:"timestamp"`
}
