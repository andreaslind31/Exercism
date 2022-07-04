package chance

import (
	"fmt"
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
	time := rand.Int63()
	fmt.Println(time)
}

// RollADie returns a random int d with 1 <= d <= 20
func RollADie() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 20

	die, _ := fmt.Println(rand.Intn((max - min) + min))

	return die
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	return rand.Float64()
}

// ShuffleAnimals returns a slice with all eight animal strings in random order
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

	rand.Shuffle(len(animals), func(i, j int) { animals[i], animals[j] = animals[j], animals[i] })

	return animals
}
