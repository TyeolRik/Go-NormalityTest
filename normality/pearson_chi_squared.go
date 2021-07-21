package normality

import "math"

// https://en.wikipedia.org/wiki/Pearson%27s_chi-squared_test
// https://github.com/cran/nortest/blob/master/R/pearson.test.R
func PearsonChiSquared(data *[]float64) (chiSquared float64, P_value float64) {
	n := len(*data)
	n_float64 := float64(n)

	dfd := 2.0
	mean, sd := Get_AverageAndStandardDeviation(data)

	n_classes := math.Ceil(2.0 * math.Pow(n_float64, 0.4))
	num := make([]float64, n)
	for i := range num {
		num[i] = math.Floor(1.0 + n_classes*NormalDistribution_CDF((*data)[i], mean, sd))
	}
	count := tabulate(&num, int(n_classes))
	xpec := make([]float64, int(n_classes))
	for i := range xpec {
		xpec[i] = n_float64 / n_classes
	}
	P_value = 0.0
	for i := range count {
		P_value += (float64(count[i]) - xpec[i]) * (float64(count[i]) - xpec[i]) / xpec[i]
	}
	chiSquared = P_value
	P_value = pchisq(P_value, n_classes-dfd-1, 1)
	return
}
