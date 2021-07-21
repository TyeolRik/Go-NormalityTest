package normality

import (
	"fmt"
	"os"
	"sort"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var testStatistics, P_value float64
var testNumber int = 1
var totalPassFail int = 0
var PassFail string = "PASS"

var t table.Writer = table.NewWriter()

func DoNormalityTest(data *[]float64) {
	sort.Float64s(*data)

	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Test Name", "Test Statistics", "P-value", "Pass/Fail"})

	// Step 1. D'Agostino's K-squared test
	testStatistics, P_value = D_AgostinosKsquared(data)
	doTest("D'Agostino's K-squared test", testStatistics, P_value, 0.05)
	t.AppendSeparator()

	// Step 2. Jarque–Bera test
	testStatistics, P_value = JarqueBera(data)
	doTest("Jarque–Bera test", testStatistics, P_value, 0.05)
	t.AppendSeparator()

	// Step 3. Anderson–Darling test
	testStatistics, P_value = AndersonDarling(*data)
	doTest("Anderson–Darling test", testStatistics, P_value, 0.05)
	t.AppendSeparator()

	// Step 4. Cramér–von Mises criterion
	testStatistics, P_value = CramerVonMises(*data)
	doTest("Cramér–von Mises criterion", testStatistics, P_value, 0.05)
	t.AppendSeparator()

	// Step 5. Lilliefors test (same as Kolmogorov–Smirnov test)
	testStatistics, P_value = Lilliefors(*data)
	doTest("Lilliefors test", testStatistics, P_value, 0.05)
	t.AppendSeparator()

	// Step 6. Shapiro–Francia test (same as Shapiro–Wilk test)
	testStatistics, P_value = ShapiroFrancia(*data)
	doTest("Shapiro–Francia test", testStatistics, P_value, 0.05)
	t.AppendSeparator()

	// Step 7. Pearson's chi-squared test
	testStatistics, P_value = PearsonChiSquared(data)
	doTest("Pearson's chi-squared test", testStatistics, P_value, 0.05)

	var resultColor text.Colors = text.Colors{text.FgHiRed}
	resultPassFail := "FAIL"
	if totalPassFail > (testNumber / 2) {
		resultPassFail = "PASS"
		resultColor = text.Colors{text.FgHiGreen}
	}

	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 3, Align: text.AlignRight, AlignHeader: text.AlignCenter, AlignFooter: text.AlignCenter},
		{Number: 4, Align: text.AlignRight, AlignHeader: text.AlignCenter, AlignFooter: text.AlignCenter, ColorsFooter: resultColor},
		{Number: 5, Align: text.AlignCenter, AlignHeader: text.AlignCenter, AlignFooter: text.AlignCenter, ColorsFooter: resultColor},
	})

	t.AppendFooter(table.Row{"", "", "", "Result", resultPassFail})
	t.Render()

	QQplot(data)
}

func doTest(testName string, testStatistics float64, P_value float64, significanceLevel float64) {
	if P_value > 0.05 {
		totalPassFail++
		PassFail = "PASS"
	} else {
		PassFail = "Fail"
	}
	t.AppendRow([]interface{}{testNumber, testName, fmt.Sprintf("%.08f", testStatistics), fmt.Sprintf("%.08f", P_value), PassFail})
	testNumber++
}
