package router

import (
	"fmt"
	"hd_wallet_go/wallet/handlers"

	"github.com/aws/aws-lambda-go/events"
)

type Router struct{}

func (r *Router) Route(req events.APIGatewayProxyRequest) (string, error) {
	if req.Path == "/wallet/generate" && req.RequestContext.HTTPMethod == "GET" {
		return handlers.Generate()
	} else if req.Path == "/wallet/restore" && req.RequestContext.HTTPMethod == "POST" {
		return handlers.Restore()
	} else {
		return "", fmt.Errorf("%s match not found", req.Path)
	}
}