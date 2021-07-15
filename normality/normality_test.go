package normality_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/tyeolrik/Go-NormalityTest/normality"
)

func TestQQPlot(t *testing.T) {
	// data := normality.ReadFileAsFloat64Slice("../sample_data/QELP_DataSet_057.txt")
	data := normality.ReadFileAsFloat64Slice("../sample_data/QQplotDataset.txt")
	normality.QQplot(data)
}

func TestHistogram(t *testing.T) {
	data := normality.ReadFileAsFloat64Slice("../sample_data/QQplotDataset.txt")
	normality.Histogram(data)
}

func TestNormSINV(t *testing.T) {
	fmt.Println(normality.Norm_S_INV(0.75))
}

// Same Example from Official Microsoft document
func TestGet_StandardDeviation(t *testing.T) {
	data := []float64{1345, 1301, 1368, 1322, 1310, 1370, 1318, 1350, 1303, 1299}
	mean, stdev_s := normality.Get_AverageAndStandardDeviation(data)
	// Expected values (Correct answer)
	// mean = 1328.6
	// stdev_s = 27.46391571984349
	fmt.Println("              Mean : ", mean)
	fmt.Println("Standard Deviation : ", stdev_s)
}

func TestGetQuantileType7(t *testing.T) {
	data := []float64{1345, 1301, 1368, 1322, 1310, 1370, 1318, 1350, 1303, 1299}
	sort.Float64s(data)
	Q1position := normality.GetQuantileType7(&data, 0.25)
	Q3position := normality.GetQuantileType7(&data, 0.75)
	fmt.Println("Q1", Q1position)
	fmt.Println("Q3", Q3position)
}
