// Package weather do some weather calculations.
package weather

var (
    // CurrentCondition is condition.
	CurrentCondition string
    // CurrentLocation is location.
	CurrentLocation  string
)

// Forecast returns some.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
