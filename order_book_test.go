package orderbook

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	orderBook := NewOrderBook(BTC_USD)

	if bidSize := orderBook.GetBidSize(Price(999)); bidSize != 0 {
		t.Errorf("BID size should be 0 for an empty orderbook")
	}

	if askSize := orderBook.GetAskSize(Price(999)); askSize != 0 {
		t.Errorf("ASK size should be 0 for an empty orderbook")
	}

	orderBook.Add(BUY, Price(999), Volume(100))
	orderBook.Add(SELL, Price(1001), Volume(200))

	if bidSize := orderBook.GetBidSize(Price(999)); bidSize != 100 {
		t.Errorf("BID size is invalid. got %v expected %v", bidSize, 100)
	}

	if askSize := orderBook.GetAskSize(Price(1001)); askSize != 200 {
		t.Errorf("ASK size is invalid. got %v expected %v", askSize, 200)
	}

	expectedResult := []level{
		level{
			bidPrice: 999,
			bidSize:  100,
			askPrice: 1001,
			askSize:  200,
		},
	}
	results := levels(orderBook)
	if !reflect.DeepEqual(results, expectedResult) {
		t.Errorf("Values don't match: got %v expected %v", results, expectedResult)
	}
}
func TestAdd_IncreaseBidVolume(t *testing.T) {
	orderBook := NewOrderBook(BTC_USD)

	orderBook.Add(BUY, Price(999), Volume(100))
	orderBook.Add(BUY, Price(999), Volume(200))

	if bidSize := orderBook.GetBidSize(Price(999)); bidSize != 300 {
		t.Errorf("BID size is invalid. got %v expected %v", bidSize, 300)
	}

	if askSize := orderBook.GetAskSize(Price(1001)); askSize != 0 {
		t.Errorf("ASK size is invalid. got %v expected %v", askSize, 0)
	}

	expectedResult := []level{
		level{
			bidPrice: 999,
			bidSize:  300,
			askPrice: 0,
			askSize:  0,
		},
	}
	results := levels(orderBook)
	if !reflect.DeepEqual(results, expectedResult) {
		t.Errorf("Values don't match: got %v expected %v", results, expectedResult)
	}
}
func TestBestBidAndAsk(t *testing.T) {
	orderBook := NewOrderBook(BTC_USD)

	if k, _ := orderBook.GetBestBid(); k != 0 {
		t.Errorf("Best bid price is invalid. Got %v expected %v", k, 0)
	}

	if k, _ := orderBook.GetBestAsk(); k != 0 {
		t.Errorf("Best ask price is invalid. Got %v expected %v", k, 0)
	}

	orderBook.Add(BUY, Price(999), Volume(100))
	orderBook.Add(SELL, Price(1051), Volume(200))
	orderBook.Add(BUY, Price(2000), Volume(50))
	orderBook.Add(SELL, Price(1050), Volume(200))

	if k, _ := orderBook.GetBestBid(); k != 2000 {
		t.Errorf("Best bid price is invalid. Got %v expected %v", k, 2000)
	}

	if k, _ := orderBook.GetBestAsk(); k != 1050 {
		t.Errorf("Best ask price is invalid. Got %v expected %v", k, 1050)
	}
}
