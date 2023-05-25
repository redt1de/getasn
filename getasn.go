package getasn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type IPInfo struct {
	IP       string `yaml:"ip"`
	Hostname string `yaml:"hostname"`
	City     string `yaml:"city"`
	Region   string `yaml:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
	Org      string `json:"org"`
	Zip      string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

func GetASN(ip string) (*IPInfo, error) {
	ret := IPInfo{}
	resp, err := http.Get("https://ipinfo.io/" + ip)
	if err != nil {
		return &ret, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &ret, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return &ret, err
	}
	return &ret, nil
}
