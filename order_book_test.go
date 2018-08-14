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
	orderBook.Add(SELL, Price(1001), Volume(200))

	assert.Equal(t, Volume(100), orderBook.GetBidSize(Price(999)))

	assert.Equal(t, Volume(200), orderBook.GetAskSize(Price(1001)))

	expectedResult := []level{
		level{
			bidPrice: 999,
			bidSize:  100,
			askPrice: 1001,
			askSize:  200,
		},
	}
	assert.Equal(t, expectedResult, levels(orderBook))
}
func TestAdd_IncreaseBidVolume(t *testing.T) {
	orderBook := NewOrderBook(BTC_USD)

	orderBook.Add(BUY, Price(999), Volume(100))
	orderBook.Add(BUY, Price(999), Volume(200))

	assert.Equal(t, Volume(300), orderBook.GetBidSize(Price(999)))
	assert.Equal(t, Volume(0), orderBook.GetAskSize(Price(1001)))

	expectedResult := []level{
		level{
			bidPrice: 999,
			bidSize:  300,
			askPrice: 0,
			askSize:  0,
		},
	}
	assert.Equal(t, expectedResult, levels(orderBook))
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
