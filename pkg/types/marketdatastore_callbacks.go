// Code generated by "callbackgen -type MarketDataStore"; DO NOT EDIT.

package types

import ()

func (store *MarketDataStore) OnKLineWindowUpdate(cb func(interval Interval, klines KLineWindow)) {
	store.kLineWindowUpdateCallbacks = append(store.kLineWindowUpdateCallbacks, cb)
}

func (store *MarketDataStore) EmitKLineWindowUpdate(interval Interval, klines KLineWindow) {
	for _, cb := range store.kLineWindowUpdateCallbacks {
		cb(interval, klines)
	}
}

func (store *MarketDataStore) OnKLineClosed(cb func(k KLine)) {
	store.kLineClosedCallbacks = append(store.kLineClosedCallbacks, cb)
}

func (store *MarketDataStore) EmitKLineClosed(k KLine) {
	for _, cb := range store.kLineClosedCallbacks {
		cb(k)
	}
}
