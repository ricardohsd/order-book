package orderbook

// Side is a set of types that defines which side the order was executed.
type Side int

const (
	BUY Side = 1 + iota
	SELL
)

// Price describes the price of a market order.
// To facilitate usage we avoid using float.
// Instead we must store prices as unsigned integers by multiplying it by 100000000.
//
// Example: 1.00000061 should be stored as 100000061
type Price uint64

// Volume describes the bid/ask quantity of a market order.
// To facilitate usage we avoid using float.
// Instead we must store prices as unsigned integers by multiplying it by 100000000.
//
// Example: 1.00000061 should be stored as 100000061
type Volume uint64

// Instrument describes which markets exists.
type Instrument string

var (
	BTC_USD Instrument = "btc-usd"
	BTC_ETH Instrument = "btc-eth"
)
