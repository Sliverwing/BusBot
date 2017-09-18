package models

type LineHttpResponce struct {
	Status Status     `json:"status"`
	Result LineResult `json:"result"`
}

type LineSearchHttpResponse struct {
	Status Status       `json:"status"`
	Result []LineResult `json:"result"`
}

type LineResult struct {
	ID               string    `json:"string"`
	Area             int       `json:"area"`
	LineName         string    `json:"lineName"`
	StartStationName string    `json:"startStationName"`
	EndStationName   string    `json:"endStationName"`
	Stations         []Station `json:"stations"`
}

type Station struct {
	ID          string  `json:"id"`
	Area        int     `json:"area"`
	StationName string  `json:"stationName"`
	Lng         float32 `json:"lng"`
	Lat         float32 `json:"lat"`
	Buslines    string  `json:"buslines"`
	State       string  `json:"state"`
	UpdateTime  string  `json:"updateTime"`
	// busLineList ignored
}
