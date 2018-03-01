// // //    $ gst-launch-1.0 -v pulsesrc ! audioconvert ! audioresample ! audio/x-raw,channels=1,rate=16000 ! filesink location=/dev/stdout | livecaption
package recognize

import (
	"fmt"
	"log"

	speech "cloud.google.com/go/speech/apiv1"
	"golang.org/x/net/context"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func Recognize(data []byte) string {
	ctx := context.Background()

	// [START init]
	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// [END init]

	// [START request]
	//data, err := ioutil.ReadFile("audio.raw")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Send the contents of the audio file with the encoding and
	// and sample rate information to be transcripted.
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz: 44100,
			LanguageCode:    "cs-CZ",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	})
	// [END request]
	fmt.Println(resp.Results[0].Alternatives[0].Transcript, err)

	return resp.Results[0].Alternatives[0].Transcript
}

// // // //    $ gst-launch-1.0 -v pulsesrc ! audioconvert ! audioresample ! audio/x-raw,channels=1,rate=16000 ! filesink location=/dev/stdout | livecaption
// package recognize
//
// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
//
// 	speech "cloud.google.com/go/speech/apiv1"
// 	"golang.org/x/net/context"
// 	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
// )
//
// func Recognize() {
// 	ctx := context.Background()
//
// 	// [START speech_streaming_mic_recognize]
// 	client, err := speech.NewClient(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	stream, err := client.StreamingRecognize(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Send the initial configuration message.
// 	if err := stream.Send(&speechpb.StreamingRecognizeRequest{
// 		StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
// 			StreamingConfig: &speechpb.StreamingRecognitionConfig{
// 				Config: &speechpb.RecognitionConfig{
// 					Encoding:        speechpb.RecognitionConfig_LINEAR16,
// 					SampleRateHertz: 16000,
// 					LanguageCode:    "en-US",
// 				},
// 			},
// 		},
// 	}); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	go func() {
// 		// Pipe stdin to the API.
// 		buf := make([]byte, 1024)
// 		for {
// 			n, err := os.Stdin.Read(buf)
// 			if err == io.EOF {
// 				// Nothing else to pipe, close the stream.
// 				if err := stream.CloseSend(); err != nil {
// 					log.Fatalf("Could not close stream: %v", err)
// 				}
// 				return
// 			}
// 			if err != nil {
// 				log.Printf("Could not read from stdin: %v", err)
// 				continue
// 			}
// 			if err = stream.Send(&speechpb.StreamingRecognizeRequest{
// 				StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
// 					AudioContent: buf[:n],
// 				},
// 			}); err != nil {
// 				log.Printf("Could not send audio: %v", err)
// 			}
// 		}
// 	}()
//
// 	for {
// 		resp, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("Cannot stream results: %v", err)
// 		}
// 		if err := resp.Error; err != nil {
// 			log.Fatalf("Could not recognize: %v", err)
// 		}
// 		for _, result := range resp.Results {
// 			fmt.Printf("Result: %+v\n", result)
// 		}
// 	}
// 	// [END speech_streaming_mic_recognize]
// }
