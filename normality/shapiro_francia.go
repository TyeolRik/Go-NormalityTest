package normality

// https://en.wikipedia.org/wiki/Shapiro-Francia_test
// Compared to the Shapiro–Wilk test, Shapiro–Francia test statistic is easier to compute
// In practice the Shapiro–Wilk and Shapiro–Francia variants are about equally good.

// https://github.com/cran/nortest/blob/master/R/sf.test.R
func ShapiroFrancia(data *[]float64) (testStatistics float64, P_value float64) {

}
