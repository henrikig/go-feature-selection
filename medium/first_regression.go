package medium

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/sajari/regression"
)

func firstRegression() {
	f, err := os.Open("./datasets/training.csv")

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

	var r regression.Regression

	r.SetObserved("Price")
	r.SetVar(0, "Grade")

	for _, record := range records[1:] {
		price, err := strconv.ParseFloat(record[2], 64)

		if err != nil {
			log.Fatal(err)
		}

		grade, err := strconv.ParseFloat(record[11], 64)

		if err != nil {
			log.Fatal(err)
		}

		r.Train(regression.DataPoint(price, []float64{grade}))
	}

	start := time.Now()
	r.Run()
	elapsed := time.Since(start)
	fmt.Printf("Training took %v\n", elapsed)

	fmt.Printf("\nRegression Formula:\n%v\n\n", r.Formula)
	fmt.Printf("\nR2:\n%v\n\n", math.Sqrt(r.R2))

}
