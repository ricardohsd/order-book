package orderbook

// level provides a struct for transversing both orderbook's treemaps
type level struct {
	bidPrice Price
	bidSize  Volume
	askPrice Price
	askSize  Volume
}

// Iterates through bids and asks and returns each bid leveled up with an ask in the same position.
func levels(orderBook *OrderBook) []level {
	var levels []level
	var bids []Price
	var asks []Price

	itb := orderBook.bids.Iterator()
	for itb.Next() {
		key := itb.Key().(Price)
		bids = append(bids, key)
	}

	ita := orderBook.asks.Iterator()
	for ita.Next() {
		key := ita.Key().(Price)
		asks = append(asks, key)
	}

	maxItems := 0

	if len(bids) > len(asks) {
		maxItems = len(bids)
	} else {
		maxItems = len(asks)
	}

	for i := 0; i < maxItems; i++ {
		bidPrice := Price(0)
		askPrice := Price(0)

		if bids == nil || i >= len(bids) {
			bidPrice = Price(0)
		} else {
			bidPrice = bids[i]
		}

		if asks == nil || i >= len(asks) {
			askPrice = Price(0)
		} else {
			askPrice = asks[i]
		}

		bidSize := orderBook.GetBidSize(bidPrice)
		askSize := orderBook.GetAskSize(askPrice)

		levels = append(levels, level{
			bidPrice: bidPrice,
			bidSize:  bidSize,
			askPrice: askPrice,
			askSize:  askSize,
		})
	}

	return levels
}
