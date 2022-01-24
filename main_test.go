package main

import (
	"reflect"
	"testing"
)

func TestReform(t *testing.T) {

	body := []byte(`{
		"ip": "1.1.1.1",
		"is_eu": false,
		"city": null,
		"region": null,
		"region_code": null,
		"country_name": "Australia",
		"country_code": "AU",
		"continent_name": "Oceania",
		"continent_code": "OC",
		"latitude": -33.494,
		"longitude": 143.2104,
		"postal": null,
		"calling_code": "61",
		"flag": "https://ipdata.co/flags/au.png",
		"emoji_flag": "\ud83c\udde6\ud83c\uddfa",
		"emoji_unicode": "U+1F1E6 U+1F1FA",
		"asn": {
			"asn": "AS13335",
			"name": "Cloudflare, Inc.",
			"domain": "cloudflare.com",
			"route": "1.1.1.0/24",
			"type": "business"
		},
		"languages": [
			{
				"name": "English",
				"native": "English",
				"code": "en"
			}
		],
		"currency": {
			"name": "Australian Dollar",
			"code": "AUD",
			"symbol": "AU$",
			"native": "$",
			"plural": "Australian dollars"
		},
		"time_zone": {
			"name": "Australia/Sydney",
			"abbr": "AEDT",
			"offset": "+1100",
			"is_dst": true,
			"current_time": "2022-01-24T13:57:15+11:00"
		},
		"threat": {
			"is_tor": false,
			"is_proxy": false,
			"is_anonymous": false,
			"is_known_attacker": false,
			"is_known_abuser": false,
			"is_threat": false,
			"is_bogon": false
		},
		"count": "12"
	}`)

	outTest := &Output{
		IP:          "1.1.1.1",
		CountryName: "Australia",
		TZOffset:    "+1100",
		IsProxy:     false,
	}

	out, _ := reform(body)

	if !reflect.DeepEqual(out, outTest) {
		t.Fatalf("expected %v, got: %v", outTest, out)
	}

}
