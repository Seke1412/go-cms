package main

import (
	"net/http"
	"path/filepath"
	Models "sample/go-cms/models"
	"time"

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

	router.Static("/upload-images", "./upload-images")

	router.GET("/samples", getSamples)

	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/playground-upload", uploadFile)
	router.POST("/sample-upload", uploadForm)

	router.Run("localhost:8080")
}

func uploadForm(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	files := form.File["files"]
	articleTitle := c.PostForm("articleTitle")
	articleContent := c.PostForm("articleContent")
	photoArrayAsString := ""

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		relativeFilePath := "upload-images/" + filename
		photoArrayAsString += relativeFilePath + ","
		// TODO: check file exist before upload
		if err := c.SaveUploadedFile(file, "./"+relativeFilePath); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}
	}

	newSample := Models.Sample{
		Title:   articleTitle,
		Content: articleContent,
		Photo:   photoArrayAsString,
	}

	newSample.Create()

	c.String(http.StatusOK, "Success!")
}

func uploadFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	files := form.File["files"]

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		// TODO: check file exist before upload
		if err := c.SaveUploadedFile(file, "./upload-images/"+filename); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}
	}

	c.String(http.StatusOK,
		"Uploaded successfully %d files with fields name=%s and email=%s.", len(files))
}

func getSamples(c *gin.Context) {
	ss, err := Models.GetSamples()
	if err == nil {
		c.IndentedJSON(http.StatusOK, ss)
	} else {
		panic(err)
	}

}
