// Package weather sets the weather in a specific location.
package weather

// CurrentCondition represents the weather condition as a string.
var CurrentCondition string

// CurrentLocation represents the city as a string.
var CurrentLocation string

// Forecast returns an aggregated string of CurrentLocation and CurrentCondition which sets by the paramaters city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
