package medium

import (
	"image/color"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func firstRegressionPlot() {
	f, err := os.Open("./datasets/house_data.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	df := dataframe.ReadCSV(f)

	pts := make(plotter.XYs, df.Nrow())

	ptsPred := make(plotter.XYs, df.Nrow())

	yVals := df.Col("price").Float()

	for i, floatVal := range df.Col("grade").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
		ptsPred[i].X = floatVal
		ptsPred[i].Y = predict(floatVal)

	}

	p := plot.New()

	p.X.Label.Text = "grade"
	p.Y.Label.Text = "house price"

	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(pts)

	if err != nil {
		log.Fatal(err)
	}

	s.GlyphStyle.Radius = vg.Points(2)
	s.GlyphStyle.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}

	l, err := plotter.NewLine(ptsPred)

	if err != nil {
		log.Fatal(err)
	}

	l.LineStyle.Width = vg.Points(0.5)
	l.LineStyle.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
	l.LineStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}

	p.Add(s, l)

	if err := p.Save(10*vg.Centimeter, 10*vg.Centimeter, "./graphs/first_regression.png"); err != nil {
		log.Fatal(err)
	}
}
