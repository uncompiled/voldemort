package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// swapCmd represents the swap command
var swapCmd = &cobra.Command{
	Use:   "swap",
	Short: "Face Swap",
	Long:  `If a match is found in the image, do a faceswap`,
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
	RootCmd.AddCommand(swapCmd)
	swapCmd.Flags().StringVarP(&url, "image", "i", "", "image to identify")
}

func swap(imageURL string) {

}
