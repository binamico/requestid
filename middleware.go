package requestid

import (
	"net/http"

	"github.com/google/uuid"
)

// InjectRequestIDToCtx добавляет идентификатор запроса в контекст.
// Генерирует новый, если в заголовке не было значения.
func InjectRequestIDToCtx(header string) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			requestID := req.Header.Get(header)
			if requestID == "" {
				requestID = uuid.New().String()
			}
			ctx := InjectRequestID(req.Context(), requestID)
			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}
