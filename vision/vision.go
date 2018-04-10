package vision

import (
	"context"
	"fmt"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

func VisionInit() {
	_ = context.Background()
	_ = vision.ImageAnnotatorClient{}
	_ = os.Open
}

func DetectURL(url string) string {
	ctx := context.Background()

	fmt.Println(url)

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		panic(err)
	}

	image := vision.NewImageFromURI(url)
	annotations, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		panic(err)
	}

	response := ""

	if len(annotations) == 0 {
		response = "No labels found."
	} else {
		response = "Labels:"
		for _, annotation := range annotations {
			response += annotation.Description
		}
	}
	return response
}
