package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

var inputData string

// loadTrainingData reads in the JSON training data into a []TrainingSet
func loadTrainingData(inputFile string) []TrainingData {
	input, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	var trainingSet []TrainingData
	err = json.Unmarshal(input, &trainingSet)
	if err != nil {
		panic(err)
	}

	return trainingSet
}

// processTrainingData sends individual TrainingData to facebox
func processTrainingData(trainingSet []TrainingData) {
	for index := range trainingSet {
		data, err := json.Marshal(trainingSet[index])
		if err != nil {
			panic(err)
		}

		statusCode := sendTrainingData(string(data))
		if statusCode != 200 {
			fmt.Printf("%d: %s\n", statusCode, trainingSet[index].URL)
		} else {
			fmt.Println("Success!")
		}
	}
}

// sendTrainingData sends the HTTP requests
func sendTrainingData(data string) int {
	request := gorequest.New()
	resp, _, _ := request.Post("http://localhost:8080/facebox/teach").
		Send(data).
		End()
	return resp.StatusCode
}

// trainCmd represents the train command
var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "Teaches our facebox how to recognize Voldemort",
	Long:  `Uses the output from the data step to train our facebox`,
	Run: func(cmd *cobra.Command, args []string) {
		trainingSet := loadTrainingData(inputData)
		processTrainingData(trainingSet)
	},
}

func init() {
	RootCmd.AddCommand(trainCmd)

	trainCmd.Flags().StringVarP(&inputData, "input", "i", "training/data.json", "training data")
}
