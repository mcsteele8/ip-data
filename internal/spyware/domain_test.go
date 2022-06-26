package spyware

import (
	"io"
	"ip-data/tools/mock"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestSpyware_getDomainInfo(t *testing.T) {
	type fields struct {
		httpClient HttpClient
	}
	type args struct {
		domain string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:   "success_run",
			fields: fields{httpClient: &mock.ClientMock{Response: &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(domainMockResponse))}}},
			args:   args{domain: "cyderes.com"},
			want: &DomainApiResponse{
				Domain:      "cyderes.com",
				DomainID:    "2177835482_DOMAIN_COM-VRSN",
				Status:      "clientDeleteProhibited https://icann.org/epp#clientDeleteProhibited",
				CreateDate:  parseTime("2017-10-23T15:51:21Z"),
				UpdateDate:  parseTime("2022-06-09T15:58:55Z"),
				ExpireDate:  parseTime("2022-10-23T15:51:21Z"),
				DomainAge:   1707,
				WhoisServer: "whois.godaddy.com",
			},
			wantErr: false,
		},
		{
			name:    "fail_run",
			fields:  fields{httpClient: &mock.ClientMock{Response: &http.Response{StatusCode: http.StatusBadRequest, Body: io.NopCloser(strings.NewReader(""))}}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Spyware{
				httpClient: tt.fields.httpClient,
			}
			got, err := s.getDomainInfo(tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("Spyware.getDomainInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.(*DomainApiResponse).Domain != tt.want.(*DomainApiResponse).Domain && got.(*DomainApiResponse).DomainID != tt.want.(*DomainApiResponse).DomainID {
				t.Errorf("Spyware.getDomainInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func parseTime(timeString string) time.Time {
	t, _ := time.Parse(time.RFC3339, timeString)
	return t
}

var domainMockResponse = `
{
    "domain": "cyderes.com",
    "domain_id": "2177835482_DOMAIN_COM-VRSN",
    "status": "clientDeleteProhibited https://icann.org/epp#clientDeleteProhibited",
    "create_date": "2017-10-23T15:51:21Z",
    "update_date": "2022-06-09T15:58:55Z",
    "expire_date": "2022-10-23T15:51:21Z",
    "domain_age": 1707,
    "whois_server": "whois.godaddy.com",
    "registrar": {
        "iana_id": "",
        "name": "",
        "url": ""
    },
    "registrant": {
        "name": "",
        "organization": "",
        "street_address": "",
        "city": "",
        "region": "",
        "zip_code": "",
        "country": "",
        "phone": "",
        "fax": "",
        "email": ""
    },
    "admin": {
        "name": "",
        "organization": "",
        "street_address": "",
        "city": "",
        "region": "",
        "zip_code": "",
        "country": "",
        "phone": "",
        "fax": "",
        "email": ""
    },
    "tech": {
        "name": "",
        "organization": "",
        "street_address": "",
        "city": "",
        "region": "",
        "zip_code": "",
        "country": "",
        "phone": "",
        "fax": "",
        "email": ""
    },
    "billing": {
        "name": "",
        "organization": "",
        "street_address": "",
        "city": "",
        "region": "",
        "zip_code": "",
        "country": "",
        "phone": "",
        "fax": "",
        "email": ""
    },
    "nameservers": [
        "NS03.DOMAINCONTROL.COM"
    ]
}
`
