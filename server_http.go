package mygrpc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

// NewHTTPServer makes a new mygrpc HTTP service.
func NewHTTPServer(ctx context.Context, endpoints MyEndpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/lastName", httptransport.NewServer(
		endpoints.GetLastNameEndpoint,
		decodeLastNameRequest,
		encodeLastNameResponse,
	))
	return m
}

func NewSimpleLoggingMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(lastNameRequest)
			log.Println("lastNamRequest::FirstName:", req.FirstName)

			return next(ctx, request)
		}
	}
}
