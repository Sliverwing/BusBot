package utils

import (
	"fmt"

	"github.com/levigross/grequests"
	"github.com/sliverwing/BusBot/models"
)

// DailBusDetail : Request for the detail of buses
func DailBusDetail(lineID int) (*models.BusesHttpResponce, error) {

	bus := &models.BusesHttpResponce{}
	resp, err := grequests.Get(fmt.Sprintf("http://60.210.101.86:8980/server-ue2/rest/buses/busline/370300/%d", lineID), GetOptions())
	if err != nil {
		return nil, err
	}
	resp.JSON(&bus)
	return bus, nil
}
