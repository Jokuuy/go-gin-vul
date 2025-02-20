package unAuth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

// Bad case
func GetImage(c *gin.Context) {

	var url string
	// Check the request method
	if c.Request.Method == "GET" {
		url = c.Query("q")
	} else if c.Request.Method == "POST" {
		url = c.PostForm("q")
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"err": err,
		})
		fmt.Println("get image failed")
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
	c.JSON(200, gin.H{
		"success": string(body),
	})
}

func GetImageSafe(c *gin.Context) {

	var q string
	// Check the request method
	if c.Request.Method == "GET" {
		q = c.Query("q")
	} else if c.Request.Method == "POST" {
		q = c.PostForm("q")
	}

	url := "https://test.image.com/path/?q="
	res, err := http.Get(url + q)
	if err != nil {
		fmt.Println(err)
		fmt.Println("get image failed")
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
	c.JSON(200, gin.H{
		"success": string(body),
	})
}
