package utils

import (
	"github.com/levigross/grequests"
)

// GetOptions can retuen instance of options for every requests
func GetOptions() *grequests.RequestOptions {
	return &grequests.RequestOptions{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:46.0) Gecko/20100101 Firefox/46.0",
	}
}
