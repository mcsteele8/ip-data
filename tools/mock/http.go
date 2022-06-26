package mock

import "net/http"

type ClientMock struct {
	Response *http.Response
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	return c.Response, nil
}
