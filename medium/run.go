package medium

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
)

func run() {

	file, err := os.Open("./datasets/house_data.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data := csv.NewReader(file)
	data.FieldsPerRecord = 21

	records, err := data.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	gradeRecords := make([][]string, len(records))

	for i := range records {
		record := make([]string, 2)
		record[0], record[1] = records[i][11], records[i][2]
		gradeRecords[i] = record
	}

	shuffled := make([][]string, len(gradeRecords)-1)
	perm := rand.Perm(len(records) - 1)

	for i, v := range perm {
		shuffled[v] = gradeRecords[i+1]
	}

	trainingIdx := len(shuffled) * 4 / 5
	trainingSet := shuffled[1 : trainingIdx+1]

	testingSet := shuffled[trainingIdx+1:]

	sets := map[string][][]string{
		"./datasets/training_short.csv": trainingSet,
		"./datasets/testing_short.csv":  testingSet,
	}

	for fn, dataset := range sets {
		f, err := os.Create(fn)

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		out := csv.NewWriter(f)

		if err := out.WriteAll(dataset); err != nil {
			log.Fatal(err)
		}

		out.Flush()

	}

}
