package traceid

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/thinkgos/http-middlewares/traceid"
)

func TestTraceID(t *testing.T) {
	tests := map[string]struct {
		traceIDHeader    string
		request          func() *http.Request
		expectedResponse string
	}{
		"Retrieves Request Id from default header": {
			"X-Request-Id",
			func() *http.Request {
				req, _ := http.NewRequestWithContext(context.TODO(), "GET", "/", nil)
				req.Header.Add("X-Request-Id", "req-123456")

				return req
			},
			"RequestID: req-123456",
		},
		"Retrieves Request Id from custom header": {
			"X-Trace-Id",
			func() *http.Request {
				req, _ := http.NewRequestWithContext(context.Background(), "GET", "/", nil)
				req.Header.Add("X-Trace-Id", "trace:abc123")

				return req
			},
			"RequestID: trace:abc123",
		},
	}

	for _, test := range tests {
		w := httptest.NewRecorder()

		r := gin.New()

		traceid.TraceIDHeader = test.traceIDHeader

		r.Use(TraceID())
		r.GET("/", func(c *gin.Context) {
			traceID := FromTraceID(c)
			response := fmt.Sprintf("RequestID: %s", traceID)
			_, _ = w.WriteString(response)
		})
		r.ServeHTTP(w, test.request())

		if w.Body.String() != test.expectedResponse {
			t.Fatalf("traceid was not the expected value")
		}
	}
}
