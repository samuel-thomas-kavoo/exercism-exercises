package lasagna
import "fmt"

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) (estimate int) {
    if avgPrepTime !=0 {
        estimate = len(layers) * avgPrepTime
        return
    } else {
        estimate = len(layers) * 2
        return
    }
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64){
    var noodlesCount, sauceCount int
    noodlesCount, sauceCount = 0, 0
    for i := 0; i < len(layers); i++ {
        if layers[i] == "sauce" { 
            sauceCount++
        } else if layers[i] == "noodles" {
            noodlesCount++
        } else{
            fmt.Println("This is a non-noodle and non-sauce element")
        }
    }
    noodles = 50 * noodlesCount
    sauce = float64(sauceCount) * 0.2
    return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient (FriendList []string, MyList []string) {
    i, j := len(FriendList), len(MyList) 
    secretIngredient := FriendList[i-1]
    MyList[j-1] = secretIngredient
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe (amounts []float64, portions int) []float64{
    // quantities := amounts (This does not work as only the slice header is copied so the original slice gets affected). Leads to this error: "ScaleRecipe permuted list (was [0.5 250 150 3 0.5], now [1.5 750 450 9 1.5]), should not alter inputs"
	// We instead do a deep copy using the make and copy built-in functions
    quantities := make([]float64, len(amounts)) // makes a new slice of the same length and type as amounts 
    copy(quantities, amounts) //copies the float64 values in amounts to the new quantities slice
    for i := 0; i < len(quantities); i++ {
        quantities[i] = (quantities[i]/float64(2)) * float64(portions)
    }
    return quantities
}
// The first steps would be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// We can then implement the function logic one by one and see
// an increasing number of tests passing as we implement more 
// functionality.
