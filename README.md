# NormalityTest
This project is testing variables whether following Normal distribution in Go-Language

## Reference

- Normality Test - Wiki
[https://en.wikipedia.org/wiki/Normality_test](https://en.wikipedia.org/wiki/Normality_test)

## Test Contents

### Frequentist tests

- [D'Agostino's K-squared test](https://en.wikipedia.org/wiki/D'Agostino's_K-squared_test)
- [Jarque–Bera test](https://en.wikipedia.org/wiki/Jarque-Bera_test)
- [Anderson–Darling test](https://en.wikipedia.org/wiki/Anderson-Darling_test)
- [Cramér–von Mises criterion](https://en.wikipedia.org/wiki/Cram%C3%A9r%E2%80%93von_Mises_criterion)
- [Lilliefors test](https://en.wikipedia.org/wiki/Lilliefors_test)
- [Shapiro–Francia test](https://en.wikipedia.org/wiki/Shapiro-Francia_test)
- [Pearson's chi-squared test](https://en.wikipedia.org/wiki/Pearson%27s_chi-squared_test)

### Graphical methods

- [QQ plot](https://en.wikipedia.org/wiki/Q%E2%80%93Q_plot)

## How to use

Note the example at ```main.go```

```go
var data []float64 = make([]float64, 1000)
// Fill the data as you want.
// data = []float64{1, 2, 3, 4, 5, 6, 7, 8, ...}
normality.DoNormalityTest(&data)
```

## Example Result

### Frequentist tests

```
+---+-----------------------------+-----------------+------------+-----------+
| # | TEST NAME                   | TEST STATISTICS |   P-VALUE  | PASS/FAIL |
+---+-----------------------------+-----------------+------------+-----------+
| 1 | D'Agostino's K-squared test |      4.16196344 | 0.12480763 |    PASS   |
+---+-----------------------------+-----------------+------------+-----------+
| 2 | Jarque–Bera test            |      4.21356030 | 0.12162896 |    PASS   |
+---+-----------------------------+-----------------+------------+-----------+
| 3 | Anderson–Darling test       |      0.73231813 | 0.05612021 |    PASS   |
+---+-----------------------------+-----------------+------------+-----------+
| 4 | Cramér–von Mises criterion  |      0.12945806 | 0.04465238 |    Fail   |
+---+-----------------------------+-----------------+------------+-----------+
| 5 | Lilliefors test             |      0.03248795 | 0.01464764 |    Fail   |
+---+-----------------------------+-----------------+------------+-----------+
| 6 | Shapiro–Francia test        |      0.99740260 | 0.10220217 |    PASS   |
+---+-----------------------------+-----------------+------------+-----------+
| 7 | Pearson's chi-squared test  |     40.83200000 | 0.07122503 |    PASS   |
+---+-----------------------------+-----------------+------------+-----------+
|   |                             |                 |   RESULT   |    PASS   |
+---+-----------------------------+-----------------+------------+-----------+
```

### Graphical methods

![Imgur](https://i.imgur.com/eAlcVdk.png)

Saved in ```Go-NormalityTest/output/"Q-Q plot.png"```
※ Be aware that, if all points are near of the line, called qqline, these data look like following Normal Distribution.

> If the two distributions being compared are similar, **the points in the Q–Q plot will approximately lie on the line** y = x. If the distributions are linearly related, the points in the Q–Q plot will approximately lie on a line, but not necessarily on the line y = x.
> \- [Q–Q plot, Wikipedia](https://en.wikipedia.org/wiki/Q%E2%80%93Q_plot)