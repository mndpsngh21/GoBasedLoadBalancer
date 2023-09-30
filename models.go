package main

import (
	"net/http"
	"net/http/httputil"
)

type LB struct {
	port    string
	counter int
	servers []Server
}

type Server interface {
	serveRequest(rw http.ResponseWriter, req *http.Request)
}

type httpServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}
