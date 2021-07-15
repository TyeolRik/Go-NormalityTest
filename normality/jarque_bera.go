package normality

import (
	"math"
)

// Jarque–Bera test
// https://en.wikipedia.org/wiki/Jarque-Bera_test
func JarqueBera(data *[]float64) (JB float64, P_value float64) {
	n := len(*data)
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
	S := threeSquared / math.Pow(twoSquared, 1.5)
	// Kurtosis
	K := fourSquared / (twoSquared * twoSquared)

	// k = 1, if not used in the context of regression
	// JB = [(n-k+1) / 6] * [S^2 + (0.25*(K-3)^2)]
	JB = (n_float64 / 6.0) * (S*S + 0.25*(K-3.0)*(K-3.0))
	// If the data comes from a normal distribution, the JB statistic asymptotically has a chi-squared distribution with two degrees of freedom
	P_value = GetPvalueFromChi2(JB, 2)
	return
}
