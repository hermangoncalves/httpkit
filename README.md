# httpkit

Minimal utilities to reduce boilerplate when building HTTP services in Go.  
`httpkit` is **not a framework** — it wraps the Go standard library (`net/http`) and integrates with routers like [chi](https://github.com/go-chi/chi).  

Focus: **clean developer experience, minimal API surface, no reinvention**.

---

## Case Study

When building HTTP services in Go, I often ran into the same problem: too much repetitive code just to parse JSON, handle errors, or write responses. It made APIs noisy and harder to maintain.  

To make this smoother, I created **httpkit** — a tiny library that keeps the Go standard library and routers like chi in the center, but removes the boring boilerplate. It introduces a unified handler signature, JSON helpers, and structured error handling, so developers can focus on business logic instead of wiring.  

The result is cleaner APIs, less code to maintain (roughly 30% fewer lines in typical handlers), and a developer experience that feels consistent without adding yet another heavy framework.

---

## Features

- **Unified handler signature**: `func(*httpkit.Context) error`
- **Wrapper**: `httpkit.Wrap` turns your handler into `http.HandlerFunc`
- **Error handling**: centralized `HTTPError` with JSON output
- **Helpers**: `BindJSON`, `WriteJSON`, `Error`
- **Test utilities** (planned): `NewTestContext`
- **Observability hooks** (future): `OnRequest`, `OnResponse`, `OnError`

---

## Quickstart

```go
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hermangoncalves/httpkit"
)

func main() {
	mux := chi.NewMux()
	kit := httpkit.New(mux)

	// Simple POST /hello endpoint
	mux.Post("/hello", kit.Wrap(func(c *httpkit.Context) error {
		// Parse JSON body into struct
		var body struct {
			Name string `json:"name"`
		}
		if err := c.BindJSON(&body); err != nil {
			// Any bind/validation error returns a structured HTTPError
			return err
		}

		// Return JSON response with message
		return c.JSON(http.StatusOK, httpkit.R{
			"message": "Hello, " + body.Name,
		})
	}))
	
	log.Fatal(kit.Run())
}
````

Run it:

```bash
go run main.go
curl -X POST localhost:8080/hello -d '{"name":"world"}' -H 'Content-Type: application/json'
```

---

## Philosophy

* Do not reinvent routers or middlewares — reuse `chi`, `net/http`, or other libraries.
* Reduce boilerplate for JSON, error handling, and testing.
* Keep API surface minimal and stable.
* No global configuration; everything is explicit.

---

## Roadmap

* [x] Context + HandlerFunc
* [x] Wrap
* [x] BindJSON / WriteJSON / Error
* [ ] Test helpers
* [ ] Observability hooks (future)

---

## License

MIT
