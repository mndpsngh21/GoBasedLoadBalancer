package main

// StartLoadBalancer starts the load balancer.
// It loads the configuration and starts the HTTP load balancer.
func StartLoadBalancer() {
	// Load the configuration
	config := LoadConfig()

	// Start the HTTP load balancer with the loaded configuration
	StartHttpLoadBalancer(config)
}
