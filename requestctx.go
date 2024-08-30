package requestid

import "context"

type requestIDKey struct{}

// InjectRequestIDToCtx injects the request ID in the context.
func InjectRequestIDToCtx(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, requestID)
}

// EjectRequestIDFromCtx ejects the request id from the context.
func EjectRequestIDFromCtx(ctx context.Context) (string, bool) {
	requestID, ok := ctx.Value(requestIDKey{}).(string)
	return requestID, ok
}
