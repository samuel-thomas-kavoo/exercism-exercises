// Package weather provides the tools to forecast current weather conditons in various cites in Goblinocus.
package weather
// CurrentCondition represents the current weather condition in a city as a string.
var CurrentCondition string
// CurrentLocation represent the Current Location where the weather condition is being calculated.
var CurrentLocation string
// Forecast calculates and returns a string value that represents the current location along with the current weather conditon which is also a string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
