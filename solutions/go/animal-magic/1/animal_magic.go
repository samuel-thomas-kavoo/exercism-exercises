package chance

import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	min := 1
    max := 20
    randomNumber := rand.Intn(max) + min
    return randomNumber
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
    randomNumber := rand.Float64() + float64(rand.Intn(12))
	return randomNumber
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
    sa := []string{"ant", 
"beaver",
"cat",
"dog",
"elephant",
"fox",
"giraffe",
"hedgehog"}
    rand.Shuffle(len(sa), func(i, j int){
    	sa[i], sa[j] = sa[j], sa[i]  
    })
    return sa
}
