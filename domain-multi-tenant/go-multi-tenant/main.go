package main

import (
	"fmt"
	"net/http"
)

// tenantStore (imaginary) - Simulates a map to store tenant data for simplicity
var tenantStore = map[string]string{
	"tenant1.com": "tenant-1-id",
	"tenant2.com": "tenant-2-id",
	"localhost:8080" : "local-tenant-id",
}

func main() {
	// Define a custom handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		tenantDomain := r.Host
		println("Request received for tenant: ", tenantDomain)
		tenantID, ok := tenantStore[tenantDomain]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // Handle unknown tenant
			return
		}

		// Sample protected route (access requires valid tenant)
		if r.URL.Path == "/api/data" {
			fmt.Fprintf(w, "Tenant ID: %s\n", tenantID)
			// Simulate data access based on tenantID (replace with actual logic)
			fmt.Fprintf(w, "Tenant Data (placeholder)")
		} else {
			fmt.Fprintf(w, "Invalid endpoint")
		}
	}

	// Register the handler with the default HTTP server
	http.HandleFunc("/", handler)

	// Start server on port 8080
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
