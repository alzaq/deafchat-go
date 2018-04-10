package vision

import (
	"context"

	vision "cloud.google.com/go/vision/apiv1"
)

func RecognizeURL(url string) string {
	ctx := context.Background()

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
