package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	prepTime := len(layers) * timePerLayer

	return prepTime
}

func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64

	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		} else if v == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) {
	nrOfFriendsLastIngredient := len(friendsList) - 1
	lastIngredient := friendsList[nrOfFriendsLastIngredient]

	nrOfMyLastIngredient := len(myList) - 1
	myList[nrOfMyLastIngredient] = lastIngredient
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	required := portions - 2
	scaledRecipe := make([]float64, len(quantities))

	if required == 0 {
		return quantities
	} else if required > 1 {
		for i := range quantities {
			scaledRecipe[i] += quantities[i] * float64(portions/2)
		}
	} else if required == 1 {
		for i := range quantities {
			scaledRecipe[i] += quantities[i] * (float64(portions) - 1.5)
		}
	} else {
		for i := range quantities {
			scaledRecipe[i] += quantities[i] / 2
		}
	}

	return scaledRecipe
}
