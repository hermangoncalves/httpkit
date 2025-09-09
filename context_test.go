package httpkit

import (
	"net/http/httptest"
	"testing"
)

func TestHandlerFunc(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	ctx := NewContext(rr, req)
	handler := func(c *Context) error {
		c.Writer.WriteHeader(200)
		return nil
	}

	if err := handler(ctx); err != nil {
		t.Fatalf("unexpected error ocurred: %v", err)
	}

	if rr.Code != 200 {
		t.Fatalf("expected 200, got: %v", rr.Code)

	}
}
