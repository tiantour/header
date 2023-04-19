package header

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// Response response
type Response struct{}

// NewResponse new response
func NewResponse() *Response {
	return &Response{}
}

// Set set response metadata
func (r *Response) Set(ctx context.Context, kv ...string) context.Context {
	out := metadata.Pairs(kv...)
	if in, ok := metadata.FromIncomingContext(ctx); ok {
		out = metadata.Join(out, in)
	}

	if out.Len() != 0 {
		return metadata.NewOutgoingContext(ctx, out)
	}
	return ctx
}
