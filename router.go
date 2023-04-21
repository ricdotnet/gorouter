package main

import (
	"net"
	"net/http"
)

type (
	Handler func(r *Router)

	Router struct {
		Server          *http.Server
		Listener        net.Listener
		ListenerNetwork string
		Routes          map[string]Handler
		Request         *http.Request
		Response        http.ResponseWriter
	}
)

func NewRouter() (r *Router) {
	r = &Router{
		Server:          new(http.Server),
		Routes:          make(map[string]Handler),
		ListenerNetwork: "tcp",
	}

	r.Server.Handler = r

	return
}

func (r *Router) Start(address string) /*error*/ {
	r.Server.Addr = address

	r.newServer(r.Server)

	r.Server.Serve(r.Listener)
}

func (r *Router) newServer(s *http.Server) {
	l, err := net.Listen(r.ListenerNetwork, r.Server.Addr)
	if err != nil {
		panic("something went wrong")
	}

	r.Listener = l
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Response = w
	r.Request = req

	if fn, ok := r.Routes[req.URL.Path]; ok {
		fn(r)
		return
	}

	w.WriteHeader(404)
	w.Write([]byte("Route not found"))
}

// Verb Functions
func (r *Router) Get(path string, handler Handler) {
	r.Routes[path] = handler
}
