package normality

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"

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

func DrawPlot(title string, XLabelName string, Xs []float64, YLabelName string, Ys []float64) (savedFileLocation string) {
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

	if title == "Q-Q Plot" {
		// We need QQline
		// https://github.com/wch/r-source/blob/af7f52f70101960861e5d995d3a4bec010bc89e6/src/library/stats/R/qqnorm.R#L49
		// https://stats.stackexchange.com/a/362850

		// Default Prob
		prob := [2]float64{0.25, 0.75}
		// Get 1, 3 Quantiles (0.25, 0.75)
		// 보간법 : https://mycodepia.tistory.com/18
		y := [2]float64{GetQuantileType7(&Ys, prob[0]), GetQuantileType7(&Ys, prob[1])}

		// R qnorm is same with Excel NORM_S_INV
		// https://stackoverflow.com/a/55220740/7105963
		x := [2]float64{Norm_S_INV(prob[0]), Norm_S_INV(prob[1])}

		slope := (y[1] - y[0]) / (x[1] - x[0])
		intercept := y[0] - slope*x[0]

		trendLine := plotter.NewFunction(func(f float64) float64 { return slope*f + intercept })
		trendLine.Color = color.RGBA{B: 255, A: 255}
		p.Add(trendLine)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "../output/"+title+".png"); err != nil {
		panic(err)
	}

	path, err := filepath.Abs("../output/")
	if err != nil {
		log.Fatal(err)
	}
	savedFileLocation = path + "/" + title + ".png"
	return
}

func GetAverage(data *[]float64) (average float64) {
	n_float64 := float64(len(*data))
	average = 0.0
	for _, value := range *data {
		average = average + value
	}
	average = average / n_float64
	return
}

// Same Function in Excel
// https://support.microsoft.com/en-us/office/stdev-s-function-7d69cf97-0c1f-4acf-be27-f3e83904cc23
func Get_AverageAndStandardDeviation(data *[]float64) (average float64, STDEV_S float64) {
	n_float64 := float64(len(*data))
	average = 0.0
	for _, value := range *data {
		average = average + value
	}
	average = average / n_float64

	tempSum := 0.0
	for _, value := range *data {
		tempSum = tempSum + (value-average)*(value-average)
	}
	STDEV_S = math.Sqrt(tempSum / (n_float64 - 1.0))
	return
}

func ReadFileAsFloat64Slice(fileLocation string) (data []float64) {
	data = make([]float64, 0)

	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		temp, _ := strconv.ParseFloat(scanner.Text(), 64)
		data = append(data, temp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

// Quantile in R, Type = 7
// https://github.com/wch/r-source/blob/af7f52f70101960861e5d995d3a4bec010bc89e6/src/library/stats/R/quantile.R
func GetQuantileType7(data *[]float64, prob float64) (ret float64) {
	if !sort.IsSorted(sort.Float64Slice(*data)) {
		panic("Failed to get Quantile. Not sorted! Check again")
	}
	// position := 1.0 + float64(len(*data)-1)*prob
	position := float64(len(*data)-1) * prob // R array index start from 1, But Golang starts from 0
	diff := position - float64(uint64(position))
	if diff < 0.000000001 {
		// Then prob is integer
		ret = (*data)[uint64(position)]
		return
	}
	// Need Interpolation from now on
	ret = (*data)[uint64(position)] + ((*data)[uint64(position)+1]-(*data)[uint64(position)])*(diff)
	return
}

func GetPvalueFromChi2(chi_square float64, degreeOfFreedom int) (P_value float64) {
	n := float64(degreeOfFreedom)
	P_value = pchisq(chi_square, n, 1)
	return
}

// Reference : STAT 200EF SPRING 2016, Univ. of Illinois at Urbana-Champaign
// http://courses.atlas.illinois.edu/spring2016/STAT/STAT200/pchisq.html
func gser(n float64, x float64) float64 {
	var eps = 1.e-6
	// var gln = gamnln(n)
	var gln = math.Log(math.Gamma(n / 2.0))
	var a = 0.5 * n
	var ap = a
	var sum = 1.0 / a
	var del = sum
	for n := 1; n < 101; n++ {
		ap++
		del = del * x / ap
		sum += del
		if del < sum*eps {
			break
		}
	}
	return sum * math.Exp(-x+a*math.Log(x)-gln)
}

func gcf(n float64, x float64) float64 {
	var eps = 1.e-6
	// var gln = gamnln(n)
	var gln = math.Log(math.Gamma(n / 2.0))
	var a = 0.5 * n
	var b = x + 1 - a
	var fpmin = 1.e-300
	var c = 1 / fpmin
	var d = 1 / b
	var h = d
	for i := 1; i < 101; i++ {
		var an = -float64(i) * (float64(i) - a)
		b += 2.0
		d = an*d + b
		if math.Abs(d) < fpmin {
			d = fpmin
		}
		c = b + an/c
		if math.Abs(c) < fpmin {
			c = fpmin
		}
		d = 1 / d
		var del = d * c
		h = h * del
		if math.Abs(del-1) < eps {
			break
		}
	}
	return h * math.Exp(-x+a*math.Log(x)-gln)
}

// Return the incomplete Gamma function P(n/2,x)
// Assume n is a positive integer, x>0 , won't check arguments
func gammp(n float64, x float64) float64 {
	if x < 0.5*n+1 {
		return gser(n, x)
	} else {
		return 1 - gcf(n, x)
	}
}

// Return the incomplete Gamma function Q(n/2,x)
// Assume n is a positive integer, x>0 , won't check arguments
func gammq(n float64, x float64) float64 {
	if x < 0.5*n+1 {
		return 1 - gser(n, x)
	} else {
		return gcf(n, x)
	}
}

func pchisq(chi2 float64, n float64, ptype int) float64 {
	if ptype == 1 {
		// Right tail
		return gammq(n, 0.5*chi2)
	} else {
		// Left tail
		return gammp(n, 0.5*chi2)
	}
}
