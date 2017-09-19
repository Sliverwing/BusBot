package models

import "errors"

var CachedLine *CachedLineStruct

type LineHttpResponce struct {
	Status Status     `json:"status"`
	Result LineResult `json:"result"`
}

type LineSearchHttpResponse struct {
	Status Status       `json:"status"`
	Result []LineResult `json:"result"`
}

type LineResult struct {
	ID               string    `json:"id"`
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

type CachedLineStruct struct {
	stations map[int]*map[int]Station
}

func (cl *CachedLineStruct) Init() {
	cl.stations = map[int]*map[int]Station{}
}

func (cl *CachedLineStruct) Push(id int, stations *map[int]Station) {
	cl.stations[id] = stations
}

func (cl *CachedLineStruct) IsExists(id int) bool {
	for k := range cl.stations {
		if k == id {
			return true
		}
	}
	return false
}

func (cl *CachedLineStruct) Get(id int) (*map[int]Station, error) {
	for k, v := range cl.stations {
		if k == id {
			return v, nil
		}
	}
	return nil, errors.New("Object Now Exists")
}
