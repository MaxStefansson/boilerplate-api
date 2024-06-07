package api

import (
	"log"
	"net/http"

	v1 "github.com/maxstefansson/boilerplate-api/api/v1"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	v1.SetupRoutes(router)

	middlewareChain := MiddlewareChain(
		v1.RequestLogger,
		v1.RequireAuthentication,
	)
	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(router),
	}
	log.Printf("Starting server on %s", s.addr)
	return server.ListenAndServe()
}

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next.ServeHTTP
	}
}
