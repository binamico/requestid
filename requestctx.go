package requestid

import "context"

type requestIDKey struct{}

// InjectRequestID добавляет request id в контекст
func InjectRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, requestID)
}

// EjectRequestID достает request id из запроса.
func EjectRequestID(ctx context.Context) (string, bool) {
	requestID, ok := ctx.Value(requestIDKey{}).(string)
	return requestID, ok
}
