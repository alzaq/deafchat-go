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

func main() {
	router := gin.Default()

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

		fmt.Println("1.step - data", lang, text)

		escapedText := url.PathEscape(text)

		fmt.Println("2.step - escapedText", escapedText)

		urlTTS := speech.SpeechURL(escapedText, lang)

		fmt.Println("3.step - urlTTS", urlTTS)

		response, err := http.Get(urlTTS)
		defer response.Body.Close()

		if err != nil {
			fmt.Println("Error", err)
		}

		fileName := fmt.Sprintf("sounds/%s.webm", escapedText)

		fmt.Println("4.step - fileName", fileName)

		sound, _ := os.Create(fileName)
		defer os.Remove(fileName)
		defer sound.Close()

		io.Copy(sound, response.Body)

		fmt.Println("5.step - final")

		c.File(fileName)
	})

	router.GET("/test/:lang/:text", func(c *gin.Context) {
		lang := c.Param("lang")
		text := c.Param("text")
		c.JSON(200, gin.H{
			"status": "NotFound",
			"lang":   lang,
			"text":   text,
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "NotFound",
		})
	})

	router.Run()
}
