package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"golang_poc/common"
	"golang_poc/geo-locator/service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sourceIp := req.Headers["x-forwarded-for"]

	var body []byte
	var err error
	if common.IsValidIPAddress(sourceIp, false) || common.IsValidIPAddress(sourceIp, true) {
		geoData, err := service.GetLocationData(sourceIp)
		if err != nil {
			goto errorCheck
		}
		body, err = json.Marshal(geoData)
		if err != nil {
			goto errorCheck
		}
	} else {
		body, err = json.Marshal(map[string]string{"message": "Invalid IP Address"})
	}

errorCheck:
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)

	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
