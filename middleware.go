package requestid

import (
	"net/http"

	"github.com/google/uuid"
)

// InjectRequestID adds the request ID to the context.
// Generate a new one if there was no value in the header.
// Also adds the request ID to the response headers.
// The name of the header is passed in the header parameter.
func InjectRequestID(header string) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			requestID := req.Header.Get(header)
			if requestID == "" {
				requestID = uuid.New().String()
			}
			ctx := InjectRequestIDToCtx(req.Context(), requestID)
			w.Header().Set(header, requestID)
			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}
