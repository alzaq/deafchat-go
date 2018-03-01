package speech

import "fmt"

func SpeechURL(text string, lang string) string {
	return fmt.Sprintf("https://translate.google.com/translate_tts?ie=UTF-8&tl=%s&client=tw-ob&q=%s", lang, text)
}
