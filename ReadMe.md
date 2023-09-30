# GoBasedLoadBalancer

Sample implementation for http load balancer

## Description

This project provides an example implementation of a load balancer with round-robin logic using Go. It demonstrates how to distribute incoming requests across multiple backend servers in a round-robin fashion, ensuring that each server receives an equal share of the load.

## Features

- Round-robin load balancing: Requests are distributed across multiple backend servers in a round-robin order, maintaining the order of the servers.
- HTTP reverse proxy: The load balancer acts as an HTTP reverse proxy, forwarding incoming requests to the selected backend server.
- Configuration flexibility: The backend servers can be easily configured by modifying the configuration file.

## Prerequisites

Before running the project, ensure that you have the following prerequisites:

- Go
- Configure Load Balancer listening port in servers.conf


## Configuration
The load balancer can be configured by modifying the config variable in the code. The configuration includes the following parameters:

frontend: The address and port on which the load balancer should listen for incoming requests.
backend: Comma-separated list of backend server URLs. These URLs should be in the format http://hostname:port.

# How to run
- clone it
- goto folder
- chmod +x startServers.sh
- execute shell script , it will automatically start 3 services for testing purpose and configure load balancer for it


## future scope
- add more strategies such as least-connection, sticky session
- add headers for tracking purpose
- collect statistics 