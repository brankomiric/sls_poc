package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"golang_poc/common"
)

type response struct {
	Ip string `json:"ip_address"`
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sourceIp := req.Headers["x-forwarded-for"]

	var body []byte
	var err error
	if common.IsValidIPAddress(sourceIp, "v4") || common.IsValidIPAddress(sourceIp, "v6") {
		ipAddress := &response{sourceIp}
		body, err = json.Marshal(ipAddress)
	} else {
		body, err = json.Marshal(map[string]string{"message": "Invalid IP Address"})
	}

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
