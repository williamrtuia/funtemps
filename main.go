package main

import (
	"flag"
	"fmt"
)

// Define flag variables in the main scope
var temp float64
var out string
var tempScale string

func init() {
	// Define and initialize the flag variables
	flag.Float64Var(&temp, "temp", 0.0, "temperature in degrees")
	flag.StringVar(&out, "out", "C", "output temperature scale - C - Celsius, F - Fahrenheit, K - Kelvin")
	flag.StringVar(&tempScale, "scale", "C", "input temperature scale - C - Celsius, F - Fahrenheit, K - Kelvin")
}

func main() {
	// Parse the flags
	flag.Parse()

	// Examples of using flag values
	fmt.Println(temp, out, tempScale)
	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())
	fmt.Println(isFlagPassed("out"))

	// Example of simple logic
	switch tempScale {
	case "C":
		switch out {
		case "F":
			// Call the function CelsiusToFahrenheit(temp), which should return °F
			fmt.Printf("%.2f°C is %.2f°F\n", temp, CelsiusToFahrenheit(temp))
		case "K":
			// Call the function CelsiusToKelvin(temp), which should return K
			fmt.Printf("%.2f°C is %.2fK\n", temp, CelsiusToKelvin(temp))
		default:
			// Default to Celsius
			fmt.Printf("%.2f°C\n", temp)
		}
	case "F":
		switch out {
		case "C":
			// Call the function FahrenheitToCelsius(temp), which should return °C
			fmt.Printf("%.2f°F is %.2f°C\n", temp, FahrenheitToCelsius(temp))
		case "K":
			// Call the function FahrenheitToKelvin(temp), which should return K
			fmt.Printf("%.2f°F is %.2fK\n", temp, FahrenheitToKelvin(temp))
		default:
			// Default to Fahrenheit
			fmt.Printf("%.2f°F\n", temp)
		}
	case "K":
		switch out {
		case "C":
			// Call the function KelvinToCelsius(temp), which should return °C
			fmt.Printf("%.2fK is %.2f°C\n", temp, KelvinToCelsius(temp))
		case "F":
			// Call the function KelvinToFahrenheit(temp), which should return °F
			fmt.Printf("%.2fK is %.2f°F\n", temp, KelvinToFahrenheit(temp))
		default:
			// Default to Kelvin
			fmt.Printf("%.2fK\n", temp)
		}
	default:
		// Default to Celsius
		fmt.Printf("%.2f°C\n", temp)
	}
}

// Function to check if a flag is passed
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

// Functions for temperature conversion
func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func CelsiusToKelvin(c float64) float64 {
	return c + 273.15
}

func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func FahrenheitToKelvin(f float64)

