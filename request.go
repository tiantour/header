package header

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

// Request request
type Request struct{}

// NewRequest new request
func NewRequest() *Request {
	return &Request{}
}

// Host get request host
func (r *Request) Host(ctx context.Context) string {
	return r.Get(ctx, "x-forwarded-host")
}

// IP get request ip
func (r *Request) IP(ctx context.Context) string {
	return r.Get(ctx, "x-forwarded-for")
}

// Authorization get request authorization
func (r *Request) Authorization(ctx context.Context) string {
	return r.Get(ctx, "authorization")
}

// ContentType get request content-type
func (r *Request) ContentType(ctx context.Context) string {
	return r.Get(ctx, "content-type")
}

// Tenant get request tenant
func (r *Request) Tenant(ctx context.Context) string {
	return r.Get(ctx, "tenant")
}

// UserAgent get request user-agent
func (r *Request) UserAgent(ctx context.Context) string {
	return r.Get(ctx, "user-agent")
}

// Get get request metadata
func (r *Request) Get(ctx context.Context, key string) string {
	in, ok := metadata.FromIncomingContext(ctx)
	if ok {
		gw := fmt.Sprintf("grpcgateway-%s", key)
		if value, ok := in[gw]; ok {
			return value[0]
		}

		if value, ok := in[key]; ok {
			return value[0]
		}
	}
	return ""
}
