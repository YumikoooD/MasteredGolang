package main

import "fmt"

// func main() {
// 	prices := [4]int{10, 19, 27, 35}
// 	fmt.Println(prices)

// 	newPrices := prices[1:3]
// 	fmt.Println(newPrices)

// 	newPrices = prices[1:]
// 	fmt.Println(newPrices)

// 	newPrices = prices[:3]
// 	fmt.Println(newPrices)

// 	newPrices[2] = 299
// 	fmt.Println(prices)
// we can see that prices values[2] is now 299 (I talk about PRICES not newPrices !) due to lines 14
// When you create a slices like this, you don't create a new array, it is the same array without certains element
// This is for this reason that when i modify newPrices, prices variable is also affected

// 	fmt.Println(len((newPrices)))
// 	fmt.Println(len(prices))
// }

func main() {
	prices := []int{11, 99}
	fmt.Println(prices[1])
	newPrices := append(prices, 42)
	fmt.Println(newPrices, prices) // we can see that the original array has not been modify
}
