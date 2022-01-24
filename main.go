package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

type Output struct {
	IP          string `json:"ip"`
	CountryName string `json:"country_name"`
	TZOffset    string `json:"tz_offset"`
	IsProxy     bool   `json:"is_proxy"`
}

type Response struct {
	IP            string      `json:"ip"`
	IsEu          bool        `json:"is_eu"`
	City          interface{} `json:"city"`
	Region        interface{} `json:"region"`
	RegionCode    interface{} `json:"region_code"`
	CountryName   string      `json:"country_name"`
	CountryCode   string      `json:"country_code"`
	ContinentName string      `json:"continent_name"`
	ContinentCode string      `json:"continent_code"`
	Latitude      float64     `json:"latitude"`
	Longitude     float64     `json:"longitude"`
	Postal        interface{} `json:"postal"`
	CallingCode   string      `json:"calling_code"`
	Flag          string      `json:"flag"`
	EmojiFlag     string      `json:"emoji_flag"`
	EmojiUnicode  string      `json:"emoji_unicode"`
	Asn           struct {
		Asn    string `json:"asn"`
		Name   string `json:"name"`
		Domain string `json:"domain"`
		Route  string `json:"route"`
		Type   string `json:"type"`
	} `json:"asn"`
	Languages []struct {
		Name   string `json:"name"`
		Native string `json:"native"`
		Code   string `json:"code"`
	} `json:"languages"`
	Currency struct {
		Name   string `json:"name"`
		Code   string `json:"code"`
		Symbol string `json:"symbol"`
		Native string `json:"native"`
		Plural string `json:"plural"`
	} `json:"currency"`
	TimeZone struct {
		Name        string `json:"name"`
		Abbr        string `json:"abbr"`
		Offset      string `json:"offset"`
		IsDst       bool   `json:"is_dst"`
		CurrentTime string `json:"current_time"`
	} `json:"time_zone"`
	Threat struct {
		IsTor           bool `json:"is_tor"`
		IsProxy         bool `json:"is_proxy"`
		IsAnonymous     bool `json:"is_anonymous"`
		IsKnownAttacker bool `json:"is_known_attacker"`
		IsKnownAbuser   bool `json:"is_known_abuser"`
		IsThreat        bool `json:"is_threat"`
		IsBogon         bool `json:"is_bogon"`
	} `json:"threat"`
	Count string `json:"count"`
}

func callout(ip string) ([]byte, error) {

	client := &http.Client{}

	if net.ParseIP(ip) == nil {
		return nil, errors.New(fmt.Sprint(http.StatusBadRequest))
	}

	key := os.Getenv("KEY")
	url := fmt.Sprintf("https://api.ipdata.co/%s?api-key=%s", ip, key)

	req2, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req2.Header.Add("Accept", "application/json")
	req2.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req2)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func reform(body []byte) (*Output, error) {
	var responseObject Response

	err := json.Unmarshal(body, &responseObject)
	if err != nil {
		return nil, err
	}

	out := &Output{
		IP:          responseObject.IP,
		CountryName: responseObject.CountryName,
		TZOffset:    responseObject.TimeZone.Offset,
		IsProxy:     responseObject.Threat.IsProxy,
	}

	return out, nil
}

func manager(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	body, err := callout(req.QueryStringParameters["ip"])
	if err != nil {
		return serverError(err)
	}

	out, err := reform(body)
	if err != nil {
		return serverError(err)
	}

	jsout, err := json.Marshal(out)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(jsout),
	}, nil

}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func main() {
	lambda.Start(manager)
}
