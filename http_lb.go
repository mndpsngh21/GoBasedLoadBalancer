package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

var serversCount int = 0

func (server *httpServer) serveRequest(rw http.ResponseWriter, req *http.Request) {
	server.proxy.ServeHTTP(rw, req)
}

func newHTTPServer(addr string) *httpServer {
	serverUrl, err := url.Parse(addr)
	if err != nil {
		panic("invalid url")
	}

	return &httpServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func loadServersList(config map[string]interface{}) []Server {
	data := config["backend"]
	servers, ok := data.(string)
	if !ok {
		panic("configuration invalid for servers")
	}
	fmt.Printf("Read configuration %s \n", servers)
	serverStrings := strings.Split(servers, ",")
	var serverList []Server
	for _, url := range serverStrings {
		serverList = append(serverList, newHTTPServer(url))
	}
	serversCount = len(serverList)
	return serverList
}

func (lb *LB) serveProxy(rw http.ResponseWriter, req *http.Request) {
	server := lb.servers[lb.counter]
	lb.counter = (lb.counter + 1) % len(lb.servers)
	server.serveRequest(rw, req)
}

func NewLoadBalancer(port string, servers []Server) *LB {
	return &LB{
		port:    port,
		counter: 0,
		servers: servers,
	}
}

func StartHttpLoadBalancer(config map[string]interface{}) {
	frontEnd := config["frontend"]
	frontEndConfiguration, ok := frontEnd.(string)
	if !ok {
		panic("configuration invalid for frontEnd")
	}

	port := strings.Split(frontEndConfiguration, ":")[1]

	fmt.Printf("Load balancer is listening on %s \n ", port)

	handler := NewLoadBalancer(port, loadServersList(config))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler.serveProxy(w, r)
	})
	fmt.Println(http.ListenAndServe(":"+handler.port, nil))
}
