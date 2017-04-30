package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var inputFile string
var outputFile string

// dataCmd represents the data command
var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "Generate the training data",
	Long: `Reads in a text file containing one image URL per line and 
generates data to teach the model to detect the Dark Lord	Voldemort`,
	Run: func(cmd *cobra.Command, args []string) {
		generateTrainingData(inputFile, outputFile)
		fmt.Printf("Training data exported to %s\n", outputFile)
	},
}

func init() {
	RootCmd.AddCommand(dataCmd)

	dataCmd.Flags().StringVarP(&inputFile, "input", "i", "training/images", "location of input list of images")
	dataCmd.Flags().StringVarP(&outputFile, "output", "o", "training/data.json", "location of output file")
}

func makeTrainingData(url string) TrainingData {
	return TrainingData{"Tom Riddle", url}
}

func generateTrainingData(inputFile string, outputFile string) []TrainingData {
	// Read in the list of images
	input, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// Open a JSON file
	output, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	reader := bufio.NewReader(input)
	writer := bufio.NewWriter(output)

	// Initialize an array of training data
	var trainingSet []TrainingData

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		trainingSet = append(trainingSet, makeTrainingData(string(line)))
	}

	data, err := json.Marshal(trainingSet)
	if err != nil {
		panic(err)
	}

	// Write out the data so we can check it into version control
	writer.Write(data)

	return trainingSet
}
