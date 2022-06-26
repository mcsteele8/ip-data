package spyware

import (
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Spyware struct {
	httpClient HttpClient
}

func New(httpClient *http.Client) *Spyware {
	return &Spyware{httpClient: httpClient}
}

func (s *Spyware) GetSpywareInfo(ipOrDomain string) (interface{}, error) {

	if ValidIpAddress(ipOrDomain) {
		return s.getIpInfo(ipOrDomain)
	} else {
		return s.getDomainInfo(ipOrDomain)
	}

}
