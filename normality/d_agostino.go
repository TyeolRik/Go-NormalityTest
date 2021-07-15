package normality

import (
	"log"
	"math"
)

// D'Agostino's K-squared test
// https://en.wikipedia.org/wiki/D'Agostino's_K-squared_test
func D_AgostinosKsquared(data *[]float64) (Ksquared float64, P_value float64) {
	n := len(*data)
	if n < 20 {
		log.Fatalln("Data is too small. Should be 20 or greater")
	}
	n_float64 := float64(n)

	average := GetAverage(data)
	twoSquared := 0.0
	threeSquared := 0.0
	fourSquared := 0.0
	for _, value := range *data {
		// Care for overflow
		temp := value - average
		twoSquared = twoSquared + (1.0/n_float64)*(temp*temp)
		threeSquared = threeSquared + (1.0/n_float64)*temp*temp*temp
		fourSquared = fourSquared + (1.0/n_float64)*temp*temp*temp*temp
	}
	// Skewness
	g1 := threeSquared / math.Pow(twoSquared, 1.5)
	// Kurtosis
	g2 := fourSquared/(twoSquared*twoSquared) - 3.0

	// mu1_g1 := 0.0
	mu2_g1 := (6.0 * (n_float64 - 2.0)) / ((n_float64 + 1.0) * (n_float64 + 3.0))
	// gamma1_g1 := 0.0
	gamma2_g1 := (36.0 * (n_float64 - 7.0) * (n_float64*n_float64 + 2.0*n_float64 - 5.0)) / ((n_float64 - 2.0) * (n_float64 + 5.0) * (n_float64 + 7.0) * (n_float64 + 9.0))

	mu1_g2 := -6.0 / (n_float64 + 1.0)
	mu2_g2 := (24.0 * n_float64 * (n_float64 - 2.0) * (n_float64 - 3.0)) / ((n_float64 + 1.0) * (n_float64 + 1.0) * (n_float64 + 3.0) * (n_float64 + 5.0))
	gamma1_g2 := ((6 * (n_float64*n_float64 - 5.0*n_float64 + 2.0)) / ((n_float64 + 7.0) * (n_float64 + 9.0))) * math.Sqrt((6.0*(n_float64+3.0)*(n_float64+5.0))/(n_float64*(n_float64-2.0)*(n_float64-3.0)))
	// gamma2_g2 := (36.0 * (15.0*n_float64*n_float64*n_float64*n_float64*n_float64*n_float64 - 36.0*n_float64*n_float64*n_float64*n_float64*n_float64 - 628*n_float64*n_float64*n_float64*n_float64 + 982.0*n_float64*n_float64*n_float64 + 5777.0*n_float64*n_float64 - 6402.0*n_float64 + 900.0)) / (n_float64 * (n_float64 - 3.0) * (n_float64 - 2.0) * (n_float64 + 7.0) * (n_float64 + 9.0) * (n_float64 + 11.0) * (n_float64 + 13.0))

	// Transformed sample skewness and kurtosis
	Wsquared := math.Sqrt(2*gamma2_g1+4.0) - 1.0
	delta := 1.0 / math.Sqrt(math.Log(math.Sqrt(Wsquared)))
	alphaSquared := 2.0 / (Wsquared - 1.0)
	Z1_g1 := delta * math.Asinh(g1/(math.Sqrt(alphaSquared)*math.Sqrt(mu2_g1)))

	A := 6.0 + (8.0/gamma1_g2)*((2.0/gamma1_g2)+math.Sqrt(1.0+4.0/(gamma1_g2*gamma1_g2)))
	Z2_g2 := math.Sqrt(4.5*A) * (1.0 - (2.0 / (9.0 * A)) - math.Pow((1.0-2.0/A)/(1.0+((g2-mu1_g2)/math.Sqrt(mu2_g2))*math.Sqrt(2.0/(A-4.0))), 1.0/3.0))

	Ksquared = Z1_g1*Z1_g1 + Z2_g2*Z2_g2
	P_value = GetPvalueFromChi2(Ksquared, 2)
	return
}
