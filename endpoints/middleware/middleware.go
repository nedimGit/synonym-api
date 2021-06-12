package endpoints

import (
	"net/http"

	"golang.org/x/net/context"

	env "github.com/NedimUka/sysnonyms/models/envirorment"
)

type middlewareData struct {
	Value string
}

// MiddlewareData is key in context by which user will get auth data
var MiddlewareData middlewareData

func init() {
	MiddlewareData.Value = "MiddlewareData"
}

// Post will validate if request method is POST or OPTIONS
func Post(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost && r.Method != http.MethodOptions {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		h(w, r)
	}
}

// Get will validate if request method is GET or OPTIONS
func Get(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodOptions {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		h(w, r)
	}
}

// InitMiddleware will be used to initialize the data context object
func InitMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// First check if the data has already been initialized
		rx := r.Context().Value(MiddlewareData)
		if rx != nil {
			h(w, r)
			return
		}

		// Initialize with default values. It is to be updated later on inside other middlewares and handlers.
		// Make sure that this middleware is executed first
		requestID := r.Header.Get("X-Request-Id")
		envRaw := r.Header.Get("X-Environment")
		env := env.NotSelected
		if envRaw == env.Test {
			env = env.Test
		}
		if envRaw == env.Live {
			env = env.Live
		}

		data := &env.Data{
			RequestID:   requestID,
			Environment: env,
		}
		data.AdditinalLogData = make([]interface{}, 0)

		ctx := context.WithValue(r.Context(), MiddlewareData, data)

		// Add cors
		w.Header().Set("Content-Type", "application/json charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Request-Method", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Expose-Headers", "App-Token, Status-Code, X-Environment, X-Request-Id, X-Content-Length, Content-Length")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, X-Content-Length, Content-Length, Accept-Encoding, Accept, Access-Control-Allow-Origin, Authorization, App-Token, Status-Code, Access-Control-Allow-Credentials,X-Environment,X-Request-Id,Access-Control-Request-Method")

		if r.Method == "OPTIONS" {
			return
		}

		h(w, r.WithContext(ctx))
	}
}

// Middleware is simple HandleFunc which will implement all middlewares
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Middlewares will contain all middlewares
type Middlewares struct {
	middlewares []Middleware
}

// RequestData to retrieve the Data object from the context.
// The *Data* object will be initialized if not present already
func RequestData(r *http.Request) *env.Data {
	return r.Context().Value(MiddlewareData).(*env.Data)
}

// Chain is used to call sequencially all middlewares
func Chain(middlewares ...Middleware) Middlewares {
	return Middlewares{append(([]Middleware)(nil), middlewares...)}
}

// Then is used to append additional middleware
func (m Middlewares) Then(mf http.HandlerFunc) http.HandlerFunc {
	for i := range m.middlewares {
		mf = m.middlewares[len(m.middlewares)-1-i](mf)
	}

	return mf
}

// ThenFunc is used to append additional middleware
func (m Middlewares) ThenFunc(fn http.HandlerFunc) http.HandlerFunc {
	return m.Then(fn)
}

// Append will add all middlewares into one list
func (m Middlewares) Append(middlewares ...Middleware) Middlewares {
	a := make([]Middleware, 0, len(m.middlewares)+len(middlewares))
	a = append(a, m.middlewares...)
	a = append(a, middlewares...)

	return Middlewares{a}
}

// Extend will just add additional middleware to the list of middlewares
func (m Middlewares) Extend(middlewares Middlewares) Middlewares {
	return m.Append(middlewares.middlewares...)
}
