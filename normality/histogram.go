package normality

import (
	"fmt"
	"log"
	"path/filepath"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func Histogram(data []float64) {
	p := plot.New()
	p.Title.Text = "Histogram plot"

	var value plotter.Values = data

	hist, err := plotter.NewHist(value, 20)
	if err != nil {
		panic(err)
	}
	p.Add(hist)

	if err := p.Save(3*vg.Inch, 3*vg.Inch, "../output/histogram.png"); err != nil {
		panic(err)
	}

	path, err := filepath.Abs("../output/")
	if err != nil {
		log.Fatal(err)
	}
	savedFileLocation := path + "/output/histogram.png"
	fmt.Println("File is saved at", savedFileLocation)
}
