package main

import (
	"math"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type DataPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	router.GET("/api/series1", series1Handler)
	router.GET("/api/series2", series2Handler)

	router.Run(":8080")
}

func series1Handler(c *gin.Context) {
	c.JSON(http.StatusOK, generateSineWave())
}

func series2Handler(c *gin.Context) {
	c.JSON(http.StatusOK, generateDampedOscillations())
}

// generateSineWave generates a sine wave data set.
// The sine wave is defined by the formula: y = sin(x * π / 180)
// where x is the degree (0 to 359) and y is the sine of x.
func generateSineWave() []DataPoint {
	var data []DataPoint
	for i := 0; i < 360; i++ {
		data = append(data, DataPoint{X: float64(i), Y: math.Sin(float64(i) * math.Pi / 180)})
	}
	return data
}

// generateDampedOscillations generates a set of damped oscillations.
// The damped oscillation is defined by the formula: y = e^(-x/50) * cos(x * π / 180)
// where x is the degree (0 to 359) and y represents the amplitude at x, showing exponential decay.
func generateDampedOscillations() []DataPoint {
	var data []DataPoint
	for i := 0; i < 360; i++ {
		data = append(data, DataPoint{X: float64(i), Y: math.Exp(-float64(i)/50) * math.Cos(float64(i)*math.Pi/180)})
	}
	return data
}
