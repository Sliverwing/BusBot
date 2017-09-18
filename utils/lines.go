package utils

import (
	"fmt"

	"github.com/levigross/grequests"
	"github.com/sliverwing/BusBot/models"
)

// DailLineDetail : Request for the detail of line
func DailLineDetail(id int) (*models.LineHttpResponce, error) {
	line := &models.LineHttpResponce{}
	resp, err := grequests.Get(fmt.Sprintf("http://60.210.101.86:8980/server-ue2/rest/buslines/370300/%d", id), GetOptions())
	if err != nil {
		return nil, err
	}
	resp.JSON(&line)
	return line, nil
}

func DailSearchLine(param string) (*models.LineSearchHttpResponse, error) {
	line := &models.LineSearchHttpResponse{}
	resp, err := grequests.Get(fmt.Sprintf("http://60.210.101.86:8980/server-ue2/rest/buslines/simple/370300/%s", param), GetOptions())
	if err != nil {
		return nil, err
	}
	resp.JSON(&line)
	return line, nil
}
