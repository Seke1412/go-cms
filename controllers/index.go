package controllers

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

var c = cors.New(cors.Options{
	AllowedOrigins:   []string{"http://127.0.0.1:5173", "http://localhost:5173"},
	AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Content-Type", "Authorization"},
	AllowCredentials: true,
	// Enable Debugging for testing, consider disabling in production
	Debug: true,
})

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(c)

	router.Static("/upload-images", "./upload-images")

	router.GET("/samples", getSamples)
	router.GET("/samples/:id", getSampleByID)

	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/playground-upload", uploadFile)
	router.POST("/sample-upload", uploadForm)
	router.POST("/sample-edit/:id", editSample)

	return router
}

func StartServer() {
	router := setupRouter()
	router.Run("localhost:8080")
}
