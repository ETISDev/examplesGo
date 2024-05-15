package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// tenantStore (imaginary) - Simulates a map to store tenant data for simplicity
var tenantStore = map[string]string{
	"tenant1.com": "tenant-1-id",
	"tenant2.com": "tenant-2-id",
}

func main() {
	router := gin.Default()

	// Middleware to identify tenant by domain name
	router.Use(func(c *gin.Context) {
		tenantDomain := c.Request.Host
		tenantID, ok := tenantStore[tenantDomain]
		if !ok {
			c.AbortWithStatus(http.StatusNotFound) // Handle unknown tenant
			return
		}
		c.Set("tenantID", tenantID) // Store tenantID in context
		c.Next()
	})

	// Sample protected route (access requires valid tenant)
	router.GET("/api/data", func(c *gin.Context) {
		tenantID := c.MustGet("tenantID").(string)
		fmt.Fprintf(c.Writer, "Tenant ID: %s\n", tenantID)
		// Simulate data access based on tenantID (replace with actual logic)
		fmt.Fprintf(c.Writer, "Tenant Data (placeholder)")
	})

	router.Run(":8080") // Start server on port 8080
}
