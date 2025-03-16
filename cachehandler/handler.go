package cachehandler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SetKeyValue inserts a key-value pair into the cache
func SetKeyValuePair(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad input", http.StatusBadRequest)
		return
	}

	C.SetDefault(req.Key, req.Value)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Key stored successfully")
}

// GetKeyValue retrieves a value from the cache
func GetKeyValuePair(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	value, found := C.Get(key)
	if !found {
		http.Error(w, "Key not found or expired", http.StatusNotFound)
		return
	}

	resp := map[string]interface{}{
		"key":   key,
		"value": value,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}