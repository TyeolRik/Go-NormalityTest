package normality_test

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/tyeolrik/Go-NormalityTest/normality"
)

func TestQQPlot(t *testing.T) {
	data := make([]float64, 0)

	file, err := os.Open("../sample_data/QQplotDataset.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		temp, _ := strconv.ParseFloat(scanner.Text(), 64)
		data = append(data, temp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	normality.QQplot(data)
}
