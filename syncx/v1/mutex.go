package syncx

import (
	"log"
	"sync"
	"time"
)

func NewMutex(opts ...Option) *Mutex {
	opt := makeOpt(opts)
	return &Mutex{timeout: opt.timeout, done: make(chan struct{})}
}

type Mutex struct {
	mu      sync.Mutex
	timeout time.Duration
	done    chan struct{}
}

func (m *Mutex) Lock() {
	m.mu.Lock()
	m.setTimeout()
	if m.timeout > 0 {
		go m.tick()
	}
}

func (m *Mutex) tick() {
	select {
	case <-time.After(m.timeout):
		// TODO: after timeout
		log.Println("timeout")
	case <-m.done:
	}
}

func (m *Mutex) TryLock() bool {
	if !m.mu.TryLock() {
		return false
	}
	m.setTimeout()
	if m.timeout > 0 {
		go m.tick()
	}
	return true
}

func (m *Mutex) Unlock() {
	if m.timeout > 0 {
		m.done <- struct{}{}
	}
	m.mu.Unlock()
}

func (m *Mutex) setTimeout() {
	if m.timeout == 0 && globalTimeout > 0 {
		m.timeout = globalTimeout
	}
}
