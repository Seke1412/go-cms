package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	Models "sample/go-cms/models"
	"strings"
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
	router.GET("/samples/:id", getSampleByID)

	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/playground-upload", uploadFile)
	router.POST("/sample-upload", uploadForm)
	router.POST("/sample-edit/:id", editSample)

	router.Run("localhost:8080")
}

func editSample(c *gin.Context) {
	id := c.Param("id")
	sample, err := Models.GetSample(id)
	fmt.Printf("sample %v: ", sample)

	form, err := c.MultipartForm()

	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	files := form.File["files"]
	title := c.PostForm("title")
	content := c.PostForm("content")
	remainUrls := c.PostForm("remain-urls")
	photoArrayAsString := ""

	if len(remainUrls) > 0 {
		if strings.Contains(remainUrls, ",") {
			urls := strings.Split(remainUrls, ",")
			for i := 0; i < len(urls); i++ {
				url := urls[i]
				isValidUrl := strings.Contains(sample.Photo, url)
				if isValidUrl {
					photoArrayAsString += url + ","
				}
			}
		} else {
			photoArrayAsString += remainUrls + ","
		}
	}

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		relativeFilePath := "upload-images/" + filename
		photoArrayAsString += relativeFilePath + ","
		// TODO:
		// check file exist before upload
		// sanitize file name before upload
		if err := c.SaveUploadedFile(file, "./"+relativeFilePath); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}
	}

	newSample := Models.Sample{
		Id:      id,
		Title:   title,
		Content: content,
		Photo:   photoArrayAsString,
	}

	res, err := newSample.UpdateSample()
	if err != nil {
		c.String(http.StatusBadRequest, "Cannot update sample width id %s", id)
	} else {
		fmt.Printf("res %v: ", res)
		c.String(http.StatusOK, "Success!")
	}

}

func getSampleByID(c *gin.Context) {
	id := c.Param("id")
	sample, err := Models.GetSample(id)

	if err != nil {
		c.String(http.StatusBadRequest, "Sample with id %s does not exist", id)
	} else {
		fmt.Printf("sample %v: ", sample)
		c.IndentedJSON(http.StatusOK, sample)
	}
}

func uploadForm(c *gin.Context) {
	form, err := c.MultipartForm()

	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	files := form.File["files"]
	title := c.PostForm("title")
	content := c.PostForm("content")
	photoArrayAsString := ""

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		relativeFilePath := "upload-images/" + filename
		photoArrayAsString += relativeFilePath + ","
		// TODO:
		// check file exist before upload
		// sanitize file name before upload
		if err := c.SaveUploadedFile(file, "./"+relativeFilePath); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}
	}

	photoArrayAsString = strings.TrimRight(photoArrayAsString, ",")

	newSample := Models.Sample{
		Title:   title,
		Content: content,
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
	samples, err := Models.GetSamples()
	if err == nil {
		c.IndentedJSON(http.StatusOK, samples)
	} else {
		panic(err)
	}

}
