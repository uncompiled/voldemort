package cmd

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var httpServerPort int

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves HTTP requests for face swapping",
	Long:  `Enables you to run a local HTTP server that returns face swapped images`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Listening on port %d\n", httpServerPort)

		r := mux.NewRouter()
		r.HandleFunc("/swap", FaceSwapHandler)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpServerPort), r))
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&httpServerPort, "port", "p", 8000, "port to service HTTP requests")
}

// FaceSwapHandler returns the face swapped image
func FaceSwapHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	if queryParams["image"] != nil {
		imageURL := queryParams.Get("image")

		fmt.Printf("Processing %s\n", imageURL)

		identifyResponse := identify(imageURL)
		newImage := swap(imageURL, identifyResponse)
		buf := new(bytes.Buffer)
		if err := jpeg.Encode(buf, newImage, nil); err != nil {
			panic(err)
		}
		w.Write(buf.Bytes())
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("image query parameter is required"))
	}
}
