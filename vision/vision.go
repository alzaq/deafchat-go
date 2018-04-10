package vision

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"

	vision "cloud.google.com/go/vision/apiv1"
)

func VisionInit() {
	_ = context.Background()
	_ = vision.ImageAnnotatorClient{}
	_ = os.Open
}

func DetectURL(url string) string {

	name := fmt.Sprintf("photos/%s", time.Now().Format(time.RFC3339Nano))

	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}

	defer response.Body.Close()

	//open a file for writing
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(f, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		panic(err)
	}

	f, _ = os.Open(name)
	defer os.Remove(name)
	defer f.Close()
	image, err := vision.NewImageFromReader(f)
	if err != nil {
		panic(err)
	}

	annotations, err := client.DetectLabels(ctx, image, nil, 15)
	if err != nil {
		panic(err)
	}

	text := ""

	if len(annotations) == 0 {
		text = "No labels found."
	} else {
		text = "Labels: "
		for _, annotation := range annotations {
			text += fmt.Sprintf("%s, ", annotation.Description)
		}
	}

	return text
}
