package syncx

import (
	"container/heap"
	"log"
	"sort"
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
	tm      time.Time
}

func (m *Mutex) Lock() {
	m.mu.Lock()
	if m.timeout > 0 {
		m.tick()
	}
}

func (m *Mutex) tick() {
	m.tm = time.Now().Add(m.timeout)
	if lhMu.TryLock() {
		defer lhMu.Unlock()
		heap.Push(&lh, locking{mu: m, expire: m.tm})
		return
	}
	waitingMu.Lock()
	defer waitingMu.Unlock()
	waiting = append(waiting, locking{mu: m, expire: m.tm})
}

func (m *Mutex) TryLock() bool {
	if !m.mu.TryLock() {
		return false
	}
	if m.timeout > 0 {
		m.tick()
	}
	return true
}

func (m *Mutex) Unlock() {
	if m.timeout > 0 {
		m.free()
	}
	m.mu.Unlock()
}

func (m *Mutex) free() {
	lhMu.Lock()
	defer lhMu.Unlock()
	idx, _ := sort.Find(len(lh), func(i int) int {
		if lh[i].expire.Before(m.tm) {
			return -1
		} else if lh[i].expire.After(m.tm) {
			return 1
		}
		return 0
	})
	for i := idx; i < len(lh); i++ {
		if lh[i].mu == m {
			heap.Remove(&lh, i)
			return
		}
	}
}

type locking struct {
	mu     *Mutex
	expire time.Time
}

type lockingHeap []locking

func (h *lockingHeap) Len() int { return len(*h) }

func (h *lockingHeap) Less(i, j int) bool {
	return (*h)[i].expire.Before((*h)[j].expire)
}

func (h *lockingHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *lockingHeap) Push(x any) {
	*h = append(*h, x.(locking))
}

func (h *lockingHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

var (
	lh        = make(lockingHeap, 0, 4096)
	waiting   = make([]locking, 0, 256)
	lhMu      = &sync.Mutex{}
	waitingMu = &sync.Mutex{}
)

func init() {
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			now := <-ticker.C
			var tmp []locking
			prepare := func() {
				waitingMu.Lock()
				defer waitingMu.Unlock()
				tmp = make([]locking, len(waiting))
				copy(tmp, waiting)
				waiting = waiting[:0]
			}
			work := func() {
				lhMu.Lock()
				defer lhMu.Unlock()
				for _, lock := range tmp {
					heap.Push(&lh, lock)
				}
				for len(lh) > 0 && lh[0].expire.Before(now) {
					lock := heap.Pop(&lh).(locking)
					log.Println("timeout", lock)
				}
			}
			prepare()
			work()
		}
	}()
}
