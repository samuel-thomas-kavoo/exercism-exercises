package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
    case "ace":
        return 11
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
    case "ten":
        return 10
    case "jack":
        return 10
    case "queen":
        return 10
    case "king":
        return 10
    default:
    	return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	Card1 := ParseCard(card1)
    Card2 := ParseCard(card2)
    Card3 := ParseCard(dealerCard)
    if Card1 == 11 && Card2 == 11{
        return "P"
    } else if (Card1 + Card2) == 21 && (Card3 != 11 && Card3 != 10){
        return "W"
    } else if (Card1 + Card2) == 21 && (Card3 == 11 || Card3 == 10){
        return "S"
    } else if (Card1 + Card2) >= 17 && (Card1 + Card2) <= 21{
        return "S"
    } else if (Card1 + Card2) >= 12 && (Card1 + Card2) <= 16 && Card3 < 7{
        return "S"
    } else if (Card1 + Card2) >= 12 && (Card1 + Card2) <= 16 && Card3 >= 7{
        return "H"
    } else if (Card1 + Card2) <= 11 {
        return "H"
    }
	return "S"
}
