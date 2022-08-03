package main

import (
	"net/http"
	"time"

	Models "sample/go-cms/models"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Sample struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var c = cors.New(cors.Options{
	AllowedOrigins:   []string{"http://127.0.0.1:5173", "http://localhost:5173"},
	AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Content-Type", "Authorization"},
	AllowCredentials: true,
	// Enable Debugging for testing, consider disabling in production
	Debug: true,
})

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(c)
	router.GET("/samples", getSamples)
	//router.GET("/samples/:id", getSampleByID)
	//router.POST("/samples", postSamples)

	router.Run("localhost:8080")
}

func getSamples(c *gin.Context) {
	ss, err := Models.GetSamples()
	if err == nil {
		c.IndentedJSON(http.StatusOK, ss)
	} else {
		panic(err)
	}

}

/*
func getSampleByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range samples {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "sample not found"})
}

func postSamples(c *gin.Context) {
	var newSample sample

	if err := c.BindJSON(&newSample); err != nil {
		return
	}

	// Add the new album to the slice.
	samples = append(samples, newSample)
	c.IndentedJSON(http.StatusCreated, newSample)
}
*/
