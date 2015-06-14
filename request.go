package worldtime

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	IP_ENDPOINT = "https://worldtimeiofree.p.mashape.com/ip?ipaddress="
)

var client = &http.Client{}

var api_key = ""

func SetMashapeKey(key string) {
	api_key = key
}

func Request(ip string) (r *Response, err error) {

	var ip_request *http.Request
	var ip_response *http.Response
	var ip_response_body []byte

	url := IP_ENDPOINT + ip

	if ip_request, err = http.NewRequest("GET", url, nil); err != nil {
		return
	}

	ip_request.Header.Add("Accept", "application/json")
	ip_request.Header.Add("X-Mashape-Key", api_key)

	if ip_response, err = client.Do(ip_request); err != nil {
		return
	}

	if ip_response.StatusCode != 200 {
		return nil, fmt.Errorf("Status code %d", ip_response.StatusCode)
	}

	defer ip_response.Body.Close()

	if ip_response_body, err = ioutil.ReadAll(ip_response.Body); err != nil {
		return
	}

	r, err = ParseResponse(ip_response_body)

	return
}
