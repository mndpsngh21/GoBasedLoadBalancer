#!/bin/bash

function killAllServers() {
    kill -9 $(lsof -t -i:5000)
    kill -9 $(lsof -t -i:5001)
    kill -9 $(lsof -t -i:5002)    
}

function startServers() {
killAllServers

export PORT="5000"
# Start the first server
go run sampleHttpServer/main.go &

# Wait for a moment before starting the second server
sleep 1

# Start the second server
export PORT="5001"
go run sampleHttpServer/main.go &

# Wait for a moment before starting the third server
sleep 1

# Start the third server
export PORT="5002"
go run sampleHttpServer/main.go &

sleep 2


}

startServers
echo "verify all server"

# Run the Go program
go run *.go