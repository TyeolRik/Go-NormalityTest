package normality

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Inverse standard normal cumulative distribution
// Peter J. Acklam's algorithm
// https://www.source-code.biz/snippets/vbasic/9.htm
// https://web.archive.org/web/20151030215612/http://home.online.no/~pjacklam/notes/invnorm/
func Norm_S_INV(p float64) (NormSInv float64) {
	const (
		a1 = -3.969683028665376e+01
		a2 = 2.209460984245205e+02
		a3 = -2.759285104469687e+02
		a4 = 1.383577518672690e+02
		a5 = -3.066479806614716e+01
		a6 = 2.506628277459239e+00

		b1 = -5.447609879822406e+01
		b2 = 1.615858368580409e+02
		b3 = -1.556989798598866e+02
		b4 = 6.680131188771972e+01
		b5 = -1.328068155288572e+01

		c1 = -7.784894002430293e-03
		c2 = -3.223964580411365e-01
		c3 = -2.400758277161838e+00
		c4 = -2.549732539343734e+00
		c5 = 4.374664141464968e+00
		c6 = 2.938163982698783e+00

		d1 = 7.784695709041462e-03
		d2 = 3.224671290700398e-01
		d3 = 2.445134137142996e+00
		d4 = 3.754408661907416e+00
	)
	p_low := 0.02425
	p_high := 1 - p_low
	if (p < 0) || (p > 1) {
		errorMessage := "value p: " + fmt.Sprintf("%.3f", p) + "is out-ranged."
		panic(errorMessage)
	} else if p < p_low {
		q := math.Sqrt(-2 * math.Log(p))
		NormSInv = (((((c1*q+c2)*q+c3)*q+c4)*q+c5)*q + c6) / ((((d1*q+d2)*q+d3)*q+d4)*q + 1)
	} else if p < p_high {
		q := p - 0.5
		r := q * q
		NormSInv = ((((((a1*r+a2)*r+a3)*r+a4)*r+a5)*r + a6) * q) / (((((b1*r+b2)*r+b3)*r+b4)*r+b5)*r + 1)
	} else {
		q := math.Sqrt(-2 * math.Log(1-p))
		NormSInv = -(((((c1*q+c2)*q+c3)*q+c4)*q+c5)*q + c6) / ((((d1*q+d2)*q+d3)*q+d4)*q + 1)
	}
	return
}

func DrawPlot(title string, XLabelName string, Xs []float64, YLabelName string, Ys []float64) {
	if len(Xs) != len(Ys) {
		panic("Different X, Y size!")
	}

	var n int = len(Xs)

	p := plot.New()

	p.Title.Text = title
	p.X.Label.Text = XLabelName
	p.Y.Label.Text = YLabelName
	p.Add(plotter.NewGrid())

	points := make(plotter.XYs, n)
	for i := range points {
		points[i].X = Xs[i]
		points[i].Y = Ys[i]
	}

	plotutil.AddScatters(p, points)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, title+".png"); err != nil {
		panic(err)
	}
}
