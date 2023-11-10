package main

import (
	"bytes"
	"encoding/json"
	"hd_wallet_go/wallet/router"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body []byte
	var err error

	r := router.Router{}
	result, err := r.Route(req)

	if err != nil {
		goto errorCheck
	}

	body, err = json.Marshal(map[string]interface{}{
		"message": result,
	})

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
