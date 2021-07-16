package normality

import (
	"log"
	"math"
	"sort"
)

// Anderson–Darling normality test
// Test statistic = A, P-value = P_value
// https://github.com/cran/nortest/blob/master/R/ad.test.R
func AndersonDarling(data []float64) (A float64, P_value float64) {
	n := len(data)
	n_float64 := float64(n)
	if len(data) < 8 {
		log.Fatalln("Data size should be greater than 7!")
	}
	// Need sorted data
	sort.Float64s(data)
	logp1 := make([]float64, n)
	logp2 := make([]float64, n)
	h := make([]float64, n)

	mean, sd := Get_AverageAndStandardDeviation(&data)

	for i, v := range data {
		logp1[i] = math.Log(NormalDistribution_CDF((v-mean)/sd, 0, 1))
		logp2[i] = math.Log(NormalDistribution_CDF(-(v-mean)/sd, 0, 1))
	}
	for i := range data {
		h[i] = (2.0*float64(i) + 1.0) * (logp1[i] + logp2[n-i-1])
	}

	A = -n_float64 - GetAverage(&h)
	AA := (1.0 + 0.75/n_float64 + 2.25/(n_float64*n_float64)) * A

	if AA < 0.2 {
		P_value = 1.0 - math.Exp(-13.436+101.14*AA-223.73*(AA*AA))
	} else if AA < 0.34 {
		P_value = 1.0 - math.Exp(-8.318+42.796*AA-59.938*(AA*AA))
	} else if AA < 0.6 {
		P_value = math.Exp(0.9177 - 4.279*AA - 1.38*(AA*AA))
	} else if AA < 10.0 {
		P_value = math.Exp(1.2937 - 5.709*AA + 0.0186*(AA*AA))
	} else {
		P_value = 3.7e-24
	}

	return
}
