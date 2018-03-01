package main

import (
	"deafchat-go/recognize"
	"deafchat-go/speech"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome",
		})
	})

	router.POST("/recognize/:lang", func(c *gin.Context) {
		lang := c.Param("lang")

		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}

		var data struct {
			Audio string `json:"audio"`
		}

		err = json.Unmarshal(body, &data)

		uDec, _ := base64.StdEncoding.DecodeString(data.Audio)
		value := recognize.Recognize(uDec, lang)

		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"text": value,
		})
	})

	router.GET("/speech/:lang/:text", func(c *gin.Context) {
		lang := c.Param("lang")
		text := c.Param("text")

		escapedText := url.PathEscape(text)

		urlTTS := speech.SpeechURL(escapedText, lang)

		response, err := http.Get(urlTTS)
		defer response.Body.Close()

		if err != nil {
			fmt.Println("Error", err)
		}

		fileName := fmt.Sprintf("sounds/%s.webm", escapedText)

		sound, _ := os.Create(fileName)
		defer os.Remove(fileName)
		defer sound.Close()

		io.Copy(sound, response.Body)

		c.File(fileName)
	})

	router.Run()
}
