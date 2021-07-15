package normality

import (
	"fmt"
	"sort"
)

// Inspired by Drawing QQ-plot using Excel
// https://youtu.be/nX6-j6lY9qc
func QQplot(data []float64) {
	var n int = len(data)
	var n_float64 float64 = float64(n)

	// 1. Sort Data
	sort.Float64s(data)

	// 2. Get Mean and Standard deviation
	mean, standardDeviation := Get_AverageAndStandardDeviation(data)

	ZT := make([]float64, n)
	ZA := make([]float64, n)
	for i := range ZT {
		ZT[i] = Norm_S_INV(((float64(i) + 1.0) - 0.5) / n_float64) // rank index starts 1
		ZA[i] = (data[i] - mean) / standardDeviation
	}

	plotFile := DrawPlot("Q-Q Plot", "Theoretical Quantiles", ZT, "Sample Quantiles", ZA)
	fmt.Println("File is saved at", plotFile)
}
