package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

var url string

// identifyCmd represents the identify command
var identifyCmd = &cobra.Command{
	Use:   "identify",
	Short: "Identify the faces in the photo",
	Long:  `Submits a request to detect faces trained by the model`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(url) == 0 {
			fmt.Println("Use --image to specify an image URL to identify")
		} else {
			response := identify(url)
			fmt.Println(response)
		}
	},
}

func init() {
	RootCmd.AddCommand(identifyCmd)
	identifyCmd.Flags().StringVarP(&url, "image", "i", "", "image to identify")
}

func identify(imageURL string) IdentifyResponse {
	request := gorequest.New()
	body := fmt.Sprintf(`{"url":"%s"}`, imageURL)

	resp, _, _ := request.Post("http://localhost:8080/facebox/check").
		Set("Accept", "application/json; charset=utf-8").
		Set("Content-Type", "application/json; charset=utf-8").
		Send(body).
		End()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var identifiedFaces IdentifyResponse
	err = json.Unmarshal(bodyBytes, &identifiedFaces)
	if err != nil {
		panic(err)
	}

	return identifiedFaces
}
