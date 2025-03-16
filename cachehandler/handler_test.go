package cachehandler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetKeyValuePair(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    map[string]string
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Valid request",
			requestBody: map[string]string{
				"key":   "key1",
				"value": "value1",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "Key stored successfully\n",
		},
		{
			name:           "Invalid request - missing body",
			requestBody:    nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Bad input\n",
		},
		{
			name: "Invalid request - missing key",
			requestBody: map[string]string{
				"value": "value",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Key is required\n",
		},
		{
			name: "Invalid request - missing value",
			requestBody: map[string]string{
				"key": "key",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Value is required\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitializeCache()
			var reqBody []byte
			if tt.requestBody != nil {
				reqBody, _ = json.Marshal(tt.requestBody)
			}

			req, err := http.NewRequest("POST", "/set", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(SetKeyValuePair)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.expectedBody)
			}

			C.Flush()
		})
	}
}
func TestGetKeyValuePair(t *testing.T) {
	tests := []struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedBody   string
		setupCache     func()
	}{
		{
			name:           "Valid request",
			queryParam:     "key",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"key":"key","value":"value"}` + "\n",
			setupCache: func() {
				C.SetDefault("key", "value")
			},
		},
		{
			name:           "Invalid request - missing key",
			queryParam:     "",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Key is required\n",
			setupCache:     func() {},
		},
		{
			name:           "Invalid request - key not found",
			queryParam:     "key",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Key not found or expired\n",
			setupCache:     func() {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitializeCache()
			tt.setupCache()

			req, err := http.NewRequest("GET", "/get?key="+tt.queryParam, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetKeyValuePair)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.expectedBody)
			}

			C.Flush()
		})
	}
}