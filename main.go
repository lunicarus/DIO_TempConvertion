package main

import "fmt"

func PrintBoilingPointKelvinToCelcius() {
	CelciusBoilingPoint := 100
	KelvinBoilingPoint := float32(CelciusBoilingPoint) + 273.15
	fmt.Printf("Water Boiling Point n Kelvin: %g Water Boiling Point in Celcius is %d", KelvinBoilingPoint, CelciusBoilingPoint)
}

func main() {
	PrintBoilingPointKelvinToCelcius()
}
