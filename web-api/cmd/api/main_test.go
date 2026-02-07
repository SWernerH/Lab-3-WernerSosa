package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Table-Driven Tests 

func TestHandlers(t *testing.T) {
	tests := []struct {
		name           string
		handler        http.HandlerFunc
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Home Handler",
			handler:        home,
			expectedStatus: http.StatusOK,
			expectedBody:   "Welcome to the Shapes API",
		},
		{
			name:           "Health Handler",
			handler:        health,
			expectedStatus: http.StatusOK,
			expectedBody:   "Server is running",
		},
		{
			name:           "About Handler",
			handler:        about,
			expectedStatus: http.StatusOK,
			expectedBody:   "Werner Sosa",
		},
		{
			name:           "Time Handler",
			handler:        timeHandler,
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
		{
			name:           "Quote Handler",
			handler:        quote,
			expectedStatus: http.StatusOK,
			expectedBody:   "Queen(Bohemian) 'Is this the real life? Is this just fantasy? Caught in a landslide, no escape from reality.'\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			rr := httptest.NewRecorder()

			tt.handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("got status %v, want %v", rr.Code, tt.expectedStatus)
			}

			body := rr.Body.String()

			if tt.name == "Time Handler" {
				if body == "" {
					t.Errorf("expected time handler to return a value, got empty string")
				}
			} else {
				if body != tt.expectedBody {
					t.Errorf("got body %v, want %v", body, tt.expectedBody)
				}
			}
		})
	}
}