package normality

import (
	"fmt"
	"log"
	"math"
	"sort"
)

// https://en.wikipedia.org/wiki/Shapiro-Francia_test
// Compared to the Shapiro–Wilk test, Shapiro–Francia test statistic is easier to compute
// In practice the Shapiro–Wilk and Shapiro–Francia variants are about equally good.

// https://github.com/cran/nortest/blob/master/R/sf.test.R
func ShapiroFrancia(data []float64) (testStatistics float64, P_value float64) {
	if !sort.Float64sAreSorted(data) {
		sort.Float64s(data)
	}
	n := len(data)
	n_float64 := float64(n)
	if n < 5 || n > 5000 {
		log.Fatalln("sample size must be between 5 and 5000")
	}
	y := make([]float64, n)
	temp := ppoints(n_float64, 0.375) // a = 3.0 / 8.0
	for i := range y {
		y[i] = Norm_S_INV(temp[i])
	}
	for i := 0; i < 5; i++ {
		fmt.Printf("%v : %v\n", i, y[i])
	}
	W := correlation_Pearson(&data, &y)
	W = W * W
	u := math.Log(n_float64)
	v := math.Log(u)
	mu := -1.2725 + 1.0521*(v-u)
	sig := 1.0308 - 0.26758*(v+2.0/u)
	z := (math.Log(1.0-W) - mu) / sig

	testStatistics = W
	P_value = (1.0 - NormalDistribution_CDF(z, 0, 1))
	return
}
