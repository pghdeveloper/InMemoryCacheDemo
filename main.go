package main

import (
	"fmt"
	"net/http"
	"InMemoryCacheDemo/cachehandler"
)

func main() {
	// Initialize cache before handling requests
	cachehandler.InitializeCache()

	// Set up HTTP routes
	http.HandleFunc("/set", cachehandler.SetKeyValuePair)
	http.HandleFunc("/get", cachehandler.GetKeyValuePair)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

