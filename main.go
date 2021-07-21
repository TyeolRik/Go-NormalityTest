package main

import (
	"github.com/tyeolrik/Go-NormalityTest/normality"
)

func main() {
	// PASS example
	data1 := normality.ReadFileAsFloat64Slice("./sample_data/QELP_DataSet_057.txt")
	normality.DoNormalityTest(&data1)
	data1 = nil

	// Fail example
	data2 := normality.ReadFileAsFloat64Slice("./sample_data/QQplotDataset.txt")
	normality.DoNormalityTest(&data2)
	data2 = nil
}
