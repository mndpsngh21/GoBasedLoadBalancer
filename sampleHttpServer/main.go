package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := getPortFromInput()
	addr := fmt.Sprintf(":%s", port)
	server := newHTTPServer(addr)

	http.HandleFunc("/", handleRequest)

	log.Printf("Server listening on %s", addr)
	log.Fatal(server.ListenAndServe())
}

func newHTTPServer(addr string) *http.Server {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/", handleRequest)

	return &http.Server{
		Addr:    addr,
		Handler: serverMux,
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	port := os.Getenv("PORT")
	fmt.Fprintf(w, "Hello, World! I am listening on port %s", port)
}

func getPortFromInput() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
