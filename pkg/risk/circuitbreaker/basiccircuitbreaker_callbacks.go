// Code generated by "callbackgen -type BasicCircuitBreaker"; DO NOT EDIT.

package circuitbreaker

import ()

func (b *BasicCircuitBreaker) OnPanic(cb func()) {
	b.panicCallbacks = append(b.panicCallbacks, cb)
}

func (b *BasicCircuitBreaker) EmitPanic() {
	for _, cb := range b.panicCallbacks {
		cb()
	}
}