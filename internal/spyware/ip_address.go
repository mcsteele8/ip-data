package spyware

import (
	"encoding/json"
	"io/ioutil"
	"ip-data/tools/werror"
	"net"
	"net/http"
	"strings"
)

var ipApiUrlTemplate = "http://ip-api.com/json/{ip-address}"

// ValidIpAddress() validates if input is an ip address
func ValidIpAddress(ip string) bool {
	return net.ParseIP(ip) != nil
}

// getIpInfo() queries ip-api.com to get ip information
// Link to docs https://ip-api.com/docs
func (s *Spyware) getIpInfo(ip string) (*IpApiResponse, error) {
	url := strings.ReplaceAll(ipApiUrlTemplate, "{ip-address}", ip)

	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, werror.Wrap(err, "failed to create http client")
	}

	res, err := s.httpClient.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, werror.Wrap(err, "failed to do http request")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, werror.Wrap(err, "failed to get response body")
	}

	ipInfo := IpApiResponse{}
	err = json.Unmarshal(body, &ipInfo)
	if err != nil {
		return nil, werror.Wrap(err, "failed to unmarshal response body")
	}

	return &ipInfo, nil
}

type IpApiResponse struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}
