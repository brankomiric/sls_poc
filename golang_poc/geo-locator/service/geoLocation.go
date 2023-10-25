package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type IPInfoResponse struct {
	Ip       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func GetLocationData(ipAddress string) (*IPInfoResponse, error) {
	url := fmt.Sprintf("https://ipinfo.io/%s?token=%s", ipAddress, os.Getenv("IP_INFO_TOKEN"))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var body *IPInfoResponse
	json.Unmarshal(bytes, &body)
	return body, nil
}
