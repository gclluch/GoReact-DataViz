package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestGenerateSineWave verifies the sine wave data generation.
func TestGenerateSineWave(t *testing.T) {
	data := generateSineWave()
	if len(data) != 360 {
		t.Errorf("Expected 360 data points, got %d", len(data))
	}
}

// TestGenerateDampedOscillations verifies the damped oscillations data generation.
func TestGenerateDampedOscillations(t *testing.T) {
	data := generateDampedOscillations()
	if len(data) != 360 {
		t.Errorf("Expected 360 data points, got %d", len(data))
	}
}

// TestSeries1Endpoint tests the /api/series1 endpoint.
func TestSeries1Endpoint(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/series1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	var response []DataPoint
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if len(response) != 360 {
		t.Errorf("Expected 360 data points, got %d", len(response))
	}
}

// TestSeries2Endpoint tests the /api/series2 endpoint.
func TestSeries2Endpoint(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/series2", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	var response []DataPoint
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if len(response) != 360 {
		t.Errorf("Expected 360 data points, got %d", len(response))
	}
}

// setupRouter configures the Gin router for testing.
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/api/series1", func(c *gin.Context) {
		c.JSON(http.StatusOK, generateSineWave())
	})
	router.GET("/api/series2", func(c *gin.Context) {
		c.JSON(http.StatusOK, generateDampedOscillations())
	})
	return router
}
