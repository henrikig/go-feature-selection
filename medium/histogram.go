package medium

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// histogram - plots histogram for the different columns in the house data csv file
func histogram() {
	f, err := os.Open("./datasets/house_data.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	salesData := csv.NewReader(f)

	salesData.FieldsPerRecord = 21

	records, err := salesData.ReadAll()

	header := records[0]
	records = records[1:]

	columnValues := map[int]plotter.Values{}

	for i, record := range records {
		// skip id and date columns
		for c := 2; c < salesData.FieldsPerRecord; c++ {
			if _, found := columnValues[c]; !found {
				columnValues[c] = make(plotter.Values, len(records))
			}

			floatVal, err := strconv.ParseFloat(record[c], 64)

			if err != nil {
				log.Fatal(err)
			}

			columnValues[c][i] = floatVal
		}
	}

	// draw actual graphs
	for c, values := range columnValues {
		p := plot.New()

		p.Title.Text = fmt.Sprintf("Histogram of %s", header[c])

		h, err := plotter.NewHist(values, 16)

		if err != nil {
			log.Fatal(err)
		}

		h.Normalize(1)

		p.Add(h)

		if err := p.Save(
			10*vg.Centimeter,
			10*vg.Centimeter,
			fmt.Sprintf("./graphs/%s_hist.png", header[c]),
		); err != nil {
			log.Fatal(err)
		}

	}
}
