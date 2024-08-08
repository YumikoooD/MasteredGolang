package main

import (
	"practice/prices"
)

func main() {
	taxeRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxeRates {
		pricesJob := prices.NewTaxeIncludedPriceJob(taxRate)
		pricesJob.Process()
	}
}
