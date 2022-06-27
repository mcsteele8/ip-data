package spyware

import (
	"encoding/json"
	"io/ioutil"
	"ip-data/internal/config"
	"ip-data/tools/werror"
	"net/http"
	"strings"
	"time"
)

var urlTemplate = "https://api.ip2whois.com/v2?key={your_license_key}&domain={domain_name}&format=json"

func (s *Spyware) getDomainInfo(domain string) (response DomainApiResponse, err error) {
	url := strings.ReplaceAll(urlTemplate, "{your_license_key}", config.DomainApiKey)
	url = strings.ReplaceAll(url, "{domain_name}", domain)

	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return response, werror.Wrap(err, "failed to create http client")
	}

	res, err := s.httpClient.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return response, werror.Wrap(err, "failed on domain api request")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, werror.Wrap(err, "failed to read http body")
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, werror.Wrap(err, "failed to unmarshal response body")
	}

	return response, nil
}

type DomainApiResponse struct {
	Domain      string    `json:"domain"`
	DomainID    string    `json:"domain_id"`
	Status      string    `json:"status"`
	CreateDate  time.Time `json:"create_date"`
	UpdateDate  time.Time `json:"update_date"`
	ExpireDate  time.Time `json:"expire_date"`
	DomainAge   int       `json:"domain_age"`
	WhoisServer string    `json:"whois_server"`
	Registrar   struct {
		IanaID string `json:"iana_id"`
		Name   string `json:"name"`
		URL    string `json:"url"`
	} `json:"registrar"`
	Registrant struct {
		Name          string `json:"name"`
		Organization  string `json:"organization"`
		StreetAddress string `json:"street_address"`
		City          string `json:"city"`
		Region        string `json:"region"`
		ZipCode       string `json:"zip_code"`
		Country       string `json:"country"`
		Phone         string `json:"phone"`
		Fax           string `json:"fax"`
		Email         string `json:"email"`
	} `json:"registrant"`
	Admin struct {
		Name          string `json:"name"`
		Organization  string `json:"organization"`
		StreetAddress string `json:"street_address"`
		City          string `json:"city"`
		Region        string `json:"region"`
		ZipCode       string `json:"zip_code"`
		Country       string `json:"country"`
		Phone         string `json:"phone"`
		Fax           string `json:"fax"`
		Email         string `json:"email"`
	} `json:"admin"`
	Tech struct {
		Name          string `json:"name"`
		Organization  string `json:"organization"`
		StreetAddress string `json:"street_address"`
		City          string `json:"city"`
		Region        string `json:"region"`
		ZipCode       string `json:"zip_code"`
		Country       string `json:"country"`
		Phone         string `json:"phone"`
		Fax           string `json:"fax"`
		Email         string `json:"email"`
	} `json:"tech"`
	Billing struct {
		Name          string `json:"name"`
		Organization  string `json:"organization"`
		StreetAddress string `json:"street_address"`
		City          string `json:"city"`
		Region        string `json:"region"`
		ZipCode       string `json:"zip_code"`
		Country       string `json:"country"`
		Phone         string `json:"phone"`
		Fax           string `json:"fax"`
		Email         string `json:"email"`
	} `json:"billing"`
	Nameservers []string `json:"nameservers"`
}

func (i DomainApiResponse) ToString() string {
	b, _ := json.Marshal(i)
	return string(b)
}
