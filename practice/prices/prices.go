package prices

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TaxIncludesPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludesPrices map[string]float64
}

func (job TaxIncludesPriceJob) Process() {
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	fmt.Println(result)
	getPricesFromFile("prices.txt")
}

func NewTaxeIncludedPriceJob(taxRate float64) *TaxIncludesPriceJob {
	return &TaxIncludesPriceJob{
		InputPrices: getPricesFromFile("prices.txt"),
		TaxRate:     taxRate,
	}
}

func getPricesFromFile(fileName string) []float64 {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return nil
	}

	valueText := string(data)
	listValueText := strings.Split(valueText, "\n")
	listValue := make([]float64, len(listValueText))
	for index, value := range listValueText {
		listValue[index], err = strconv.ParseFloat(value, 64)
	}

	if err != nil {
		return nil
	}

	return listValue
}
