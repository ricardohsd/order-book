package orderbook

import (
	"github.com/emirpasic/gods/maps/treemap"
)

// OrderBook represents both bids (BUY) and asks (SELL).
// For each one it will hold a height balanced binary search tree.
// The order book is a tree in which the price are keys and the values are volumes.
type OrderBook struct {
	bids       *treemap.Map // Tree to store bids (BUY) in which the key is the price and value is Volume
	asks       *treemap.Map // Tree to store asks (SELL) in which the key is the price and value is Volume
	instrument Instrument
}

func NewOrderBook(instrument Instrument) *OrderBook {
	return &OrderBook{
		bids:       treemap.NewWith(OppositePriceComparator),
		asks:       treemap.NewWith(NaturalPriceComparator),
		instrument: instrument,
	}
}

// GetBidSize returns the bid (BUY) volume size otherwise 0
func (o *OrderBook) GetBidSize(price Price) Volume {
	size, ok := o.bids.Get(price)
	if !ok {
		return Volume(0)
	}

	return size.(Volume)
}

// GetAskSize returns the ask (SELL) volume size otherwise 0
func (o *OrderBook) GetAskSize(price Price) Volume {
	size, ok := o.asks.Get(price)
	if !ok {
		return Volume(0)
	}

	return size.(Volume)
}

// Add adds a BUY or SELL order into the order orderbook according to its price.
// If the order already exists it will increase the order quantity.
func (o *OrderBook) Add(side Side, price Price, quantity Volume) {
	selectedSide := o.selectSide(side)

	var volume Volume
	volume = Volume(0)

	v, ok := selectedSide.Get(price)
	if ok {
		volume = v.(Volume)
	}

	selectedSide.Put(price, volume+quantity)
}

// GetBestBid returns the best (highest) bid and its volume
func (o *OrderBook) GetBestBid() (Price, Volume) {
	if o.bids == nil {
		return Price(0), Volume(0)
	}

	k, v := o.bids.Min()

	if k == nil {
		return Price(0), Volume(0)
	}

	return k.(Price), v.(Volume)
}

// GetBestAsk returns the best (highest) ask and its volume
func (o *OrderBook) GetBestAsk() (Price, Volume) {
	if o.asks == nil {
		return Price(0), Volume(0)
	}

	k, v := o.asks.Min()

	if k == nil {
		return Price(0), Volume(0)
	}

	return k.(Price), v.(Volume)
}

func (o *OrderBook) selectSide(side Side) *treemap.Map {
	if side == BUY {
		return o.bids
	}

	return o.asks
}
