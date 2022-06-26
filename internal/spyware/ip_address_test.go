package spyware

import (
	"io"
	"ip-data/tools/mock"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestValidIpAddress(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "good_ip_address",
			args: args{ip: "2601:680:8300:933:dd7:1e1a:c2f5:203d"},
			want: true,
		},
		{
			name: "bad_ip_address",
			args: args{ip: "1234.123.123"},
			want: false,
		},
		{
			name: "domain_name",
			args: args{ip: "google.com"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidIpAddress(tt.args.ip); got != tt.want {
				t.Errorf("ValidIpAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpyware_getIpInfo(t *testing.T) {
	type fields struct {
		httpClient HttpClient
	}
	type args struct {
		ip string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *IpApiResponse
		wantErr bool
	}{
		{
			name:   "success_run",
			fields: fields{httpClient: &mock.ClientMock{Response: &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(ipMockResponse))}}},
			args:   args{ip: "24.48.0.1"},
			want: &IpApiResponse{
				Status:      "success",
				Country:     "Canada",
				CountryCode: "CA",
				Region:      "QC",
				RegionName:  "Quebec",
				City:        "Montreal",
				Zip:         "H1K",
				Lat:         45.6085,
				Lon:         -73.5493,
				Timezone:    "America/Toronto",
				Isp:         "Le Groupe Videotron Ltee",
				Org:         "Videotron Ltee",
				As:          "AS5769 Videotron Telecom Ltee",
				Query:       "24.48.0.1",
			},
			wantErr: false,
		},
		{
			name:    "fail_run",
			fields:  fields{httpClient: &mock.ClientMock{Response: &http.Response{StatusCode: http.StatusBadRequest, Body: io.NopCloser(strings.NewReader(""))}}},
			args:    args{ip: "24.48.0.1"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Spyware{
				httpClient: tt.fields.httpClient,
			}
			got, err := s.getIpInfo(tt.args.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("Spyware.getIpInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Spyware.getIpInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

var ipMockResponse = `
{
    "status": "success",
    "country": "Canada",
    "countryCode": "CA",
    "region": "QC",
    "regionName": "Quebec",
    "city": "Montreal",
    "zip": "H1K",
    "lat": 45.6085,
    "lon": -73.5493,
    "timezone": "America/Toronto",
    "isp": "Le Groupe Videotron Ltee",
    "org": "Videotron Ltee",
    "as": "AS5769 Videotron Telecom Ltee",
    "query": "24.48.0.1"
}
`
