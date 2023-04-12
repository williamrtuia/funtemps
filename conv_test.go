package conv

import (
	"math"
	"testing"
)

func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0, want: -459.67},
		{input: 273.15, want: 32},
		{input: 373.15, want: 212},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if math.Abs(got-tc.want) > 0.0001 {
			t.Errorf("KelvinToFahrenheit(%f): expected %f, got %f", tc.input, tc.want, got)
		}
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 32, want: 273.15},
		{input: 212, want: 373.15},
		{input: -40, want: 233.15},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if math.Abs(got-tc.want) > 0.0001 {
			t.Errorf("FahrenheitToKelvin(%f): expected %f, got %f", tc.input, tc.want, got)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0, want: 273.15},
		{input: 100, want: 373.15},
		{input: -40, want: 233.15},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if math.Abs(got-tc.want) > 0.0001 {
			t.Errorf("CelsiusToKelvin(%f): expected %f, got %f", tc.input, tc.want, got)
		}
	}
}

