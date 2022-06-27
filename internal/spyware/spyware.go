package spyware

import (
	"errors"
	"ip-data/tools/werror"
	"ip-data/tools/wlog"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
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

func (s *Spyware) GetSpywareInfo(request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	wlog.New().Debugf("Request received: %v", request)

	// get the path variable from lambda proxy
	ipOrDomain, found := request.PathParameters["ip-or-domain"]
	if found {
		// path parameters are typically URL encoded so to get the value
		ipOrDomain, err = url.QueryUnescape(ipOrDomain)
		if nil != err {
			return response, werror.Wrap(err, "failed to get query ip-or-domain")
		}
	} else {
		err = errors.New("bad args. please provide an ip address or domain name")
		response.StatusCode = http.StatusBadRequest
		response.Body = err.Error()
		return
	}

	response.StatusCode = http.StatusOK
	if ValidIpAddress(ipOrDomain) {
		info, err := s.getIpInfo(ipOrDomain)
		if err != nil {
			return response, werror.Wrap(err, "failed to get ip info")
		}
		response.Body = info.ToString()
		return response, nil
	} else {
		info, err := s.getDomainInfo(ipOrDomain)
		if err != nil {
			return response, werror.Wrap(err, "failed to get ip info")
		}
		response.Body = info.ToString()
		return response, nil
	}

}
