package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, ok := units[unit]

	if ok {
		quantity, alreadyExistsInBill := bill[item]

		bill[item] = units[unit]

		if alreadyExistsInBill {
			bill[item] += quantity
		}

		return true
	}

	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, existsInBill := bill[item]
	_, existsInUnits := units[unit]

	if existsInBill && existsInUnits {
		remaining := bill[item] - units[unit]

		if remaining < 0 {
			return false
		} else if remaining == 0 {
			delete(bill, item)

			return true
		} else {
			bill[item] = remaining

			return true
		}
	}

	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]

	return value, exists
}
