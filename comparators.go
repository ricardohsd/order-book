package orderbook

// NaturalPriceComparator provides a way to compare 2 Price objects using natural order.
func NaturalPriceComparator(a, b interface{}) int {
	aAsserted := a.(Price)
	bAsserted := b.(Price)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// OppositePriceComparator provides a way to compare 2 Price objects using opposite natural order.
func OppositePriceComparator(a, b interface{}) int {
	aAsserted := a.(Price)
	bAsserted := b.(Price)
	switch {
	case aAsserted > bAsserted:
		return -1
	case aAsserted < bAsserted:
		return 1
	default:
		return 0
	}
}
