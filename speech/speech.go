package speech

import "fmt"

func SpeechURL(text string, lang string) string {
	return fmt.Sprintf("https://translate.google.com/translate_tts?ie=UTF-8&tl=%s&client=tw-ob&q=%s", lang, text)
}

//
// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/url"
// 	"os"
//
// 	"github.com/gin-gonic/gin"
// )
//
// var router *gin.Engine
//
// func main() {
// 	router = gin.Default()
// 	//router.LoadHTMLGlob("templates/*")
//
// 	router.GET("/", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "welcome",
// 		})
// 	})
//
// 	router.GET("/recognize", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"method": "recognize",
// 		})
// 	})
//
// 	router.GET("/speech/:lang/:text", func(c *gin.Context) {
// 		lang := c.Param("lang")
// 		text := c.Param("text")
//
// 		escapedText := url.PathEscape(text)
//
// 		urlTTS := SpeechURL(escapedText, lang)
//
// 		response, err := http.Get(urlTTS)
// 		defer response.Body.Close()
//
// 		if err != nil {
// 			fmt.Println("Error", err)
// 		}
//
// 		fileName := fmt.Sprintf("%s.webm", escapedText)
//
// 		sound, _ := os.Create(fileName)
// 		defer os.Remove(fileName)
// 		defer sound.Close()
//
// 		io.Copy(sound, response.Body)
// 		c.File(fileName)
// 	})
//
// 	router.Run()
// }
//
// func SpeechURL(text string, lang string) string {
// 	return fmt.Sprintf("https://translate.google.com/translate_tts?ie=UTF-8&tl=%s&client=tw-ob&q=%s", lang, text)
// }
