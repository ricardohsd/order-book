package orderbook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	orderBook := NewOrderBook(BTC_USD)

	assert.Equal(t, Volume(0), orderBook.GetBidSize(Price(999)))
	assert.Equal(t, Volume(0), orderBook.GetAskSize(Price(999)))

	orderBook.Add(BUY, Price(999), Volume(100))
	orderBook.Add(SELL, Price(1050), Volume(200))

	assert.Equal(t, []level{
		level{
			bidPrice: 999,
			bidSize:  100,
			askPrice: 1050,
			askSize:  200,
		},
	}, levels(orderBook))

	assert.Equal(t, Volume(100), orderBook.GetBidSize(Price(999)))
	assert.Equal(t, Volume(200), orderBook.GetAskSize(Price(1050)))

	orderBook.Add(BUY, Price(999), Volume(100))
	orderBook.Add(BUY, Price(999), Volume(200))

	assert.Equal(t, []level{
		level{
			bidPrice: 999,
			bidSize:  400,
			askPrice: 1050,
			askSize:  200,
		},
	}, levels(orderBook))

	assert.Equal(t, Volume(400), orderBook.GetBidSize(Price(999)))

	orderBook.Add(BUY, Price(1000), Volume(100))
	orderBook.Add(SELL, Price(1010), Volume(100))

	assert.Equal(t, []level{
		level{
			bidPrice: 1000,
			bidSize:  100,
			askPrice: 1010,
			askSize:  100,
		},
		level{
			bidPrice: 999,
			bidSize:  400,
			askPrice: 1050,
			askSize:  200,
		},
	}, levels(orderBook))
}

func TestUpdate(t *testing.T) {
	orderBook := NewOrderBook(BTC_USD)

	orderBook.Add(BUY, Price(999), Volume(100))
	orderBook.Add(SELL, Price(1050), Volume(200))

	assert.Equal(t, []level{
		level{
			bidPrice: 999,
			bidSize:  100,
			askPrice: 1050,
			askSize:  200,
		},
	}, levels(orderBook))

	assert.Equal(t, Volume(100), orderBook.GetBidSize(Price(999)))
	assert.Equal(t, Volume(200), orderBook.GetAskSize(Price(1050)))

	v := Volume(80)
	orderBook.Update(BUY, Price(999), -v)

	orderBook.Update(SELL, Price(1050), Volume(100))

	assert.Equal(t, []level{
		level{
			bidPrice: 999,
			bidSize:  20,
			askPrice: 1050,
			askSize:  300,
		},
	}, levels(orderBook))

	assert.Equal(t, Volume(20), orderBook.GetBidSize(Price(999)))
	assert.Equal(t, Volume(300), orderBook.GetAskSize(Price(1050)))

	v = Volume(300)
	orderBook.Update(SELL, Price(1050), -v)

	assert.Equal(t, []level{
		level{
			bidPrice: 999,
			bidSize:  20,
			askPrice: 0,
			askSize:  0,
		},
	}, levels(orderBook))

	assert.Equal(t, Volume(20), orderBook.GetBidSize(Price(999)))
	assert.Equal(t, Volume(0), orderBook.GetAskSize(Price(1050)))
}

func TestAdd_IncreaseBidVolume(t *testing.T) {
	orderBook := NewOrderBook(BTC_USD)

	orderBook.Add(BUY, Price(999), Volume(100))
	orderBook.Add(BUY, Price(999), Volume(200))

	assert.Equal(t, Volume(300), orderBook.GetBidSize(Price(999)))
	assert.Equal(t, Volume(0), orderBook.GetAskSize(Price(1001)))

	assert.Equal(t, []level{
		level{
			bidPrice: 999,
			bidSize:  300,
			askPrice: 0,
			askSize:  0,
		},
	}, levels(orderBook))
}
func TestBestBidAndAsk(t *testing.T) {
	orderBook := NewOrderBook(BTC_USD)

	p, v := orderBook.GetBestBid()
	assert.Equal(t, Price(0), p)
	assert.Equal(t, Volume(0), v)

	p, v = orderBook.GetBestAsk()
	assert.Equal(t, Price(0), p)
	assert.Equal(t, Volume(0), v)

	orderBook.Add(BUY, Price(999), Volume(100))
	orderBook.Add(SELL, Price(1051), Volume(200))
	orderBook.Add(BUY, Price(2000), Volume(50))
	orderBook.Add(SELL, Price(1050), Volume(200))

	p, v = orderBook.GetBestBid()
	assert.Equal(t, Price(2000), p)
	assert.Equal(t, Volume(50), v)

	p, v = orderBook.GetBestAsk()
	assert.Equal(t, Price(1050), p)
	assert.Equal(t, Volume(200), v)
}
