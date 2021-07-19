package normality

import (
	"math"
	"sort"
)

// https://github.com/cran/nortest/blob/master/R/lillie.test.R
func Lilliefors(data []float64) (testStatistics float64, P_value float64) {
	n := len(data)
	if !sort.Float64sAreSorted(data) {
		sort.Float64s(data)
	}
	n_float64 := float64(n)

	mean, variance := Get_AverageAndStandardDeviation(&data)

	DplusArr := make([]float64, n)
	DminusArr := make([]float64, n)
	for i := range DplusArr {
		DplusArr[i] = float64(i+1)/n_float64 - NormalDistribution_CDF((data[i]-mean)/variance, 0.0, 1.0)
		DminusArr[i] = NormalDistribution_CDF((data[i]-mean)/variance, 0.0, 1.0) - float64(i)/n_float64
	}
	Dplus, Dminus := DplusArr[0], DminusArr[0]
	for i := range DplusArr {
		if Dplus < DplusArr[i] {
			Dplus = DplusArr[i]
		}
		if Dminus < DminusArr[i] {
			Dminus = DminusArr[i]
		}
	}
	K := math.Max(Dplus, Dminus)
	testStatistics = K
	var Kd, nd float64
	if n <= 100 {
		Kd = K
		nd = n_float64
	} else {
		// For sample sizes m greater than 100, the same expression
		// is used with D_max, replaced by D_max * (m/100)^0.49
		// n replaced by 100.
		Kd = K * math.Pow((n_float64/100), 0.49)
		nd = 100.0
	}

	P_value = math.Exp(-7.01256*Kd*Kd*(nd+2.78019) + 2.99587*Kd*math.Sqrt(nd+2.78019) - 0.122119 + 0.974598/math.Sqrt(nd) + 1.67997/nd)

	if P_value > 0.1 {
		KK := (math.Sqrt(n_float64) - 0.01 + 0.85/math.Sqrt(n_float64)) * K
		if KK <= 0.302 {
			P_value = 1.0
		} else if KK <= 0.5 {
			P_value = 2.76773 - 19.828315*KK + 80.709644*KK*KK - 138.55152*KK*KK*KK + 81.218052*KK*KK*KK*KK
		} else if KK <= 0.9 {
			P_value = -4.901232 + 40.662806*KK - 97.490286*KK*KK + 94.029866*KK*KK*KK - 32.355711*KK*KK*KK*KK
		} else if KK <= 1.31 {
			P_value = 6.198765 - 19.558097*KK + 23.186922*KK*KK - 12.234627*KK*KK*KK + 2.423045*KK*KK*KK*KK
		} else {
			P_value = 0.0
		}
	}
	return
}
