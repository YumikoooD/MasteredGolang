package main

import "fmt"

type floatMap map[string]float64

func main() {
	website := map[string]string{
		"Google": "https://google.com",
		"Amazon": "https://amazon.com",
	}
	fmt.Println(website["Google"])
	fmt.Println(website["blabla"]) // Nothing
	website["Mikael"] = "Great Person"
	fmt.Println(website["Mikael"])
	//delete(website, "Mikael")
	newType := make(floatMap)
	newType["Mikael"] = 42.42
	newType.output(newType)
}

func (floatMap) output(o floatMap) {
	fmt.Println(o["Mikael"])
}
