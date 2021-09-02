package regression

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/linear_models"
)

func GetFitness(data *base.DenseInstances) float64 {

	trainData, testData := base.InstancesTrainTestSplit(data, 0.30)

	r := linear_models.NewLinearRegression()

	fmt.Println("Starting training...")
	start := time.Now()
	err := r.Fit(trainData)

	elapsed := time.Since(start)
	fmt.Println("Training took", elapsed)

	if err != nil {
		log.Fatal(err)
	}

	start = time.Now()
	predictions, err := r.Predict(testData)
	elapsed = time.Since(start)

	fmt.Println("Predicting took", elapsed)

	if err != nil {
		log.Fatal(err)
	}

	_, predSize := predictions.Size()

	var sum float64

	for i := 0; i < predSize; i++ {
		val, err := strconv.ParseFloat(base.GetClass(testData, i), 64)

		if err != nil {
			log.Fatal(err)
		}
		pred, err := strconv.ParseFloat(base.GetClass(predictions, i), 64)

		if err != nil {
			log.Fatal(err)
		}

		sum += math.Pow(val-pred, 2)
	}

	rmse := math.Sqrt((1 / float64(predSize)) * sum)

	return rmse

}
