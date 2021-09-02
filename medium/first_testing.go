package medium

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func firstTesting() {
	f, err := os.Open("./datasets/testing.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	salesData := csv.NewReader(f)
	salesData.FieldsPerRecord = 21

	records, err := salesData.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	records = records[1:]

	observed := make([]float64, len(records))
	predicted := make([]float64, len(records))

	var sumObserved float64

	for i, record := range records {
		price, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		observed[i] = price
		sumObserved += price

		grade, err := strconv.ParseFloat(record[11], 64)
		if err != nil {
			log.Fatal(err)
		}

		predicted[i] = predict(grade)
	}

	mean := sumObserved / float64(len(observed))

	var observedCoeff, predictCoeff float64

	for i := 0; i < len(observed); i++ {
		observedCoeff += math.Pow(observed[i]-mean, 2)
		predictCoeff += math.Pow(predicted[i]-mean, 2)
	}

	rsquared := predictCoeff / observedCoeff

	rmse := rootMeanSquaredError(predicted, observed)

	fmt.Printf("R-squared = %0.2f\n\n", rsquared)
	fmt.Printf("RMSE = %0.2f\n\n", rmse)

}

func predict(grade float64) float64 {
	return -1065201.67 + grade*209786.29
}

func rootMeanSquaredError(predicted []float64, observed []float64) float64 {
	var sum float64

	n := float64(len(predicted))

	for i := 0; i < len(predicted); i++ {
		sum += math.Pow(observed[i]-predicted[i], 2)
	}

	return math.Sqrt((1 / n) * sum)

}
