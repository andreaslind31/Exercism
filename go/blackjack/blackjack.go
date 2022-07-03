package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "king", "queen", "jack", "ten":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	firstCard := ParseCard(card1)
	secondCard := ParseCard(card2)
	dealersCard := ParseCard(dealerCard)

	handValue := firstCard + secondCard
	switch handValue {
	case 22:
		return "P"

	case 21:
		if dealersCard < 10 {
			return "W"
		}

		return "S"

	case 17, 18, 19, 20:
		return "S"

	case 12, 13, 14, 15, 16:
		if dealersCard >= 7 {
			return "H"
		}

		return "S"

	default:
		return "H"
	}
}
