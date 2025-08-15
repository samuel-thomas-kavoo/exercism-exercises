package lasagna
import "fmt"

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
    fmt.Println(OvenTime)
	fmt.Println(OvenTime - actualMinutesInOven)
    return (OvenTime - actualMinutesInOven)
}
// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	fmt.Println(2*numberOfLayers)
    return (2*numberOfLayers)
}
// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	var x int
    x = PreparationTime(numberOfLayers)
    return (x + actualMinutesInOven)
}
