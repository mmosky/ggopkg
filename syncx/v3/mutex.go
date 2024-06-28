package syncx

import (
	"time"
)

func NewMutex(opts ...Option) *Mutex {
	opt := makeOpt(opts)
	lock := make(chan struct{}, 1)
	lock <- struct{}{}
	return &Mutex{timeout: opt.timeout, lock: lock}
}

type Mutex struct {
	timeout time.Duration
	lock    chan struct{}
}

func (m *Mutex) Lock() {
	if m.lock == nil {
		panic("syncx: mutex is not initialized")
	}
	<-m.lock
}

func (m *Mutex) TryLockWithTimeout(duration time.Duration) bool {
	if m.lock == nil {
		panic("syncx: mutex is not initialized")
	}
	select {
	case <-m.lock:
		return true
	case <-time.After(duration):
		return false
	}
}

func (m *Mutex) TryLock() bool {
	if m.lock == nil {
		panic("syncx: mutex is not initialized")
	}
	select {
	case <-m.lock:
		return true
	default:
		return false
	}
}

func (m *Mutex) Unlock() {
	if m.lock == nil {
		panic("syncx: mutex is not initialized")
	}
	select {
	case m.lock <- struct{}{}:
	default:
		panic("syncx: unlock of unlocked mutex")
	}
}
