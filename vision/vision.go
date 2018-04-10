package vision

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

func VisionInit() {
	_ = context.Background()
	_ = vision.ImageAnnotatorClient{}
	_ = os.Open
}

func DetectURL(url string) string {

	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}

	defer response.Body.Close()

	//open a file for writing
	f, err := os.Create("../photos/random")
	if err != nil {
		log.Fatal(err)
	}
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(f, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		panic(err)
	}

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		panic(err)
	}
	annotations, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		panic(err)
	}

	text := ""

	if len(annotations) == 0 {
		text = "No labels found."
	} else {
		text = "Labels:"
		for _, annotation := range annotations {
			text += annotation.Description
		}
	}

	return text
}
