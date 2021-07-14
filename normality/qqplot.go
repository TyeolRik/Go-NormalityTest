package normality

import (
	"fmt"
	"math"
	"sort"
)

// Inspired by Drawing QQ-plot using Excel
// https://youtu.be/nX6-j6lY9qc
func QQplot(data []float64) {
	var n int = len(data)
	var n_float64 float64 = float64(n)

	// 1. Sort Data
	sort.Float64s(data)

	// 2-1. Get Mean
	var mean float64 = 0.0
	for _, value := range data {
		mean += value
	}
	mean = mean / n_float64
	// 2-2. Get Standard deviation
	var sumOfSquare float64 = 0.0
	for _, value := range data {
		temp := value - mean
		sumOfSquare = temp * temp
	}
	var variance float64 = sumOfSquare / (n_float64 - 1.0)
	standardDeviation := math.Sqrt(variance)

	ZT := make([]float64, n)
	ZA := make([]float64, n)
	for i := range ZT {
		ZT[i] = Norm_S_INV(((float64(i) + 1.0) - 0.5) / n_float64) // rank index starts 1
		ZA[i] = (data[i] - mean) / standardDeviation
	}

	fmt.Println("X", ZT)
	fmt.Println("Y", ZA)

	DrawPlot("Q-Q Plot", "Theoretical Quantiles", ZT, "Empirical Quantile", ZA)

}
