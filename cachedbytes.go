package main

import (
	"io"
	"sync"
)

type CachedBytes struct {
	bts  []byte
	lock sync.RWMutex
}

func NewCachedBytes(initialBytes []byte) *CachedBytes {
	cb := &CachedBytes{
		bts: initialBytes,
	}
	if cb.bts == nil {
		cb.bts = []byte{}
	}
	return cb
}

func (cb *CachedBytes) Update(newBytes []byte) {
	cb.lock.Lock()
	cb.bts = newBytes
	cb.lock.Unlock()
}

func (cb *CachedBytes) WriteTo(w io.Writer) (n int, err error) {
	cb.lock.RLock()
	defer cb.lock.RUnlock()
	return w.Write(cb.bts)
}

func (cb *CachedBytes) RLock() {
	cb.lock.RLock()
}

func (cb *CachedBytes) RUnlock() {
	cb.lock.RUnlock()
}

func (cb *CachedBytes) Length() int {
	cb.lock.RLock()
	defer cb.lock.RUnlock()
	return len(cb.bts)
}
