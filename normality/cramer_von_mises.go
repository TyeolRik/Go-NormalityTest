package normality

import (
	"log"
	"math"
	"sort"
)

// https://en.wikipedia.org/wiki/Cram%C3%A9r%E2%80%93von_Mises_criterion
// https://github.com/cran/nortest/blob/master/R/cvm.test.R
func CramerVonMises(data []float64) (W float64, P_value float64) {
	// Need sorted data
	sort.Float64s(data)
	n := len(data)
	n_float64 := float64(n)

	mean, standardDeviation := Get_AverageAndStandardDeviation(&data)
	p := make([]float64, n)
	for i := range p {
		p[i] = NormalDistribution_CDF((data[i]-mean)/standardDeviation, 0, 1) // Same with pnorm(x)
	}

	// R version
	sum := 0.0
	for i := range p {
		sum += (p[i] - (2.0*float64(i)+1.0)/(2.0*n_float64)) * (p[i] - (2.0*float64(i)+1.0)/(2.0*n_float64)) // Golang index start from 0, "not 1 like R"
	}
	W = (1.0/(12.0*n_float64) + sum)
	WW := (1.0 + 0.5/n_float64) * W
	if WW < 0.0275 {
		P_value = 1 - math.Exp(-13.953+775.5*WW-12542.61*WW*WW)
	} else if WW < 0.051 {
		P_value = 1 - math.Exp(-5.903+179.546*WW-1515.29*WW*WW)
	} else if WW < 0.092 {
		P_value = math.Exp(0.886 - 31.62*WW + 10.897*WW*WW)
	} else if WW < 1.1 {
		P_value = math.Exp(1.111 - 34.242*WW + 12.832*WW*WW)
	} else {
		log.Println("Cramér–von Mises criterion :: P-value is smaller than 7.37e-10, cannot be computed more accurately")
		P_value = 7.37e-10
		return
	}

	// Wiki Version
	// Same with above R version. Checked with some data.
	//tempSum := 0.0
	//for i := range p {
	//	tempSum += (float64(2*i+1)/(2.0*n_float64) - p[i]) * (float64(2*i+1)/(2.0*n_float64) - p[i])
	//}
	//T := 1.0/(12.0*n_float64) + tempSum

	return
}
