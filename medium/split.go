package medium

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
)

func split() {
	f, err := os.Open("./datasets/house_data.csv")

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

	header := records[0]

	shuffled := make([][]string, len(records)-1)
	perm := rand.Perm(len(records) - 1)

	for i, v := range perm {
		shuffled[v] = records[i+1]
	}

	trainingIdx := len(shuffled) * 4 / 5
	trainingSet := shuffled[1 : trainingIdx+1]

	testingSet := shuffled[trainingIdx+1:]

	sets := map[string][][]string{
		"./datasets/training.csv": trainingSet,
		"./datasets/testing.csv":  testingSet,
	}

	for fn, dataset := range sets {
		f, err := os.Create(fn)

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		out := csv.NewWriter(f)

		if err := out.Write(header); err != nil {
			log.Fatal(err)
		}

		if err := out.WriteAll(dataset); err != nil {
			log.Fatal(err)
		}

		out.Flush()

	}

}
