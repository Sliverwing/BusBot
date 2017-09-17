package models

type BusesHttpResponce struct {
	Status Status `json:"status"`
	Result []Bus  `json:"result"`
}

type Bus struct {
	BusID           string  `json:"busId"`
	Lng             float32 `json:"lng"`
	Lat             float32 `json:"lat"`
	Velocity        float32 `json:"velocity"`
	IsArrvLft       string  `json:"isArrvLft"`
	StationSeqNum   int     `json:"stationSeqNum"`
	Status          string  `json:"status"`
	BuslineID       string  `json:"buslineId"`
	ActTime         string  `json:"actTime"`
	OrgName         string  `json:"orgName"`
	AverageVelocity float32 `json:"averageVelocity"`
	Coordinate      int     `json:"coordinate"`
}
