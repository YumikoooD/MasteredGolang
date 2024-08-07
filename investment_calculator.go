package main

import "fmt"

func main() {
	yield := 10
	nbYears := 7
	finalAmount := 1000
	calculateInvestment(float32(yield), nbYears, float32(finalAmount))
}

func calculateInvestment(yield float32, nbYears int, finalAmount float32) {
	for i := 0; i != nbYears; i++ {
		finalAmount += (finalAmount * (yield / 100))
	}
	fmt.Println("After", nbYears, "years, with a", yield, "% yield, you will have", finalAmount, "dollars")
}
