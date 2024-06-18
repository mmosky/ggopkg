package syncx

import (
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	t.Run(`timeout`, func(t *testing.T) {
		m := NewMutex(WithTimeout(time.Millisecond * 200))
		m.Lock()
		time.Sleep(time.Millisecond * 240)
	})
}

func BenchmarkMutex(b *testing.B) {
	b.Run(`StdSerialLockUnlock`, func(b *testing.B) {
		mu := &sync.Mutex{}
		for i := 0; i < b.N; i++ {
			mu.Lock()
			_ = mu
			mu.Unlock()
		}
	})
	b.Run(`CompatLockUnlock`, func(b *testing.B) {
		mu := &Mutex{}
		for i := 0; i < b.N; i++ {
			mu.Lock()
			_ = mu
			mu.Unlock()
		}
	})
	b.Run(`SerialLockUnlock`, func(b *testing.B) {
		mu := NewMutex(WithTimeout(time.Millisecond * 200))
		for i := 0; i < b.N; i++ {
			mu.Lock()
			_ = mu
			mu.Unlock()
		}
	})
}
