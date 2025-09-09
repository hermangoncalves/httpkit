package httpkit

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type HandlerFunc func(c *Context) error

type HTTPError struct {
	Code    int `json:"code"`
	Message string `json:"message"`
}

func (he *HTTPError) Error() string {
	return he.Message
}

type HttpKit struct {
	mux http.Handler
}

func New(mux http.Handler) *HttpKit {
	return &HttpKit{
		mux: mux,
	}
}

func (h *HttpKit) Wrap(hf HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := NewContext(w, r)
		if err := hf(c); err != nil {
			code := http.StatusInternalServerError
			if httError, ok := err.(*HTTPError); ok {
				c.JSON(code, httError)
				return 
			}
			c.JSON(code, http.StatusText(code))
		}
	}
}

func (h *HttpKit) Run(addr ...string) error {
	address := resolveAddress(addr)

	srv := &http.Server{
		Addr:         address,
		Handler:      h.mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Printf("[HTTPKIT] 🚀 Server is running on %s\n", address)
	return srv.ListenAndServe()
}

func resolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		if port := os.Getenv("PORT"); port != "" {
			return ":" + port
		}
		return ":8080"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}
