package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	percentage := successRate / 100

	if float64(productionRate) == 0 || percentage == 0 {
		return 0
	}

	return (float64(productionRate) * percentage)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	if carsCount < 10 {
		return uint(carsCount) * 10000
	}

	groupsOfCars := carsCount / 10
	groupsCosts := groupsOfCars * 95000

	remainingCars := carsCount % 10
	if remainingCars == 0 {
		return uint(groupsCosts)
	}

	costsOfIndividualCars := remainingCars * 10000

	return uint(groupsCosts) + uint(costsOfIndividualCars)
}
