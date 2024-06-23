package fastrand

import (
	"math/rand/v2"
	"sync"
	"testing"
)

func TestUint64(t *testing.T) {
	for i := 0; i < 10; i++ {
		r := Uint64()
		t.Logf("%20d, %4d", r, r%100)
	}
}

func BenchmarkUint64(b *testing.B) {
	b.Run(`Uint64`, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Uint64()
		}
	})
	b.Run(`math/rand`, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = rand.Uint64()
		}
	})
}

func BenchmarkUint64Parallel(b *testing.B) {
	b.Run(`Uint64`, func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = Uint64()
			}
		})
	})
	b.Run(`math/rand`, func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = rand.Uint64()
			}
		})
	})
	b.Run(`Uint64_v2`, func(b *testing.B) {
		const goroutines = 10000
		var wg sync.WaitGroup
		wg.Add(goroutines)
		for i := 0; i < goroutines; i++ {
			go func() {
				defer wg.Done()
				for n := b.N; n > 0; n-- {
					_ = Uint64()
				}
			}()
		}
		wg.Wait()
	})
	b.Run(`math/rand_v2`, func(b *testing.B) {
		const goroutines = 10000
		var wg sync.WaitGroup
		wg.Add(goroutines)
		for i := 0; i < goroutines; i++ {
			go func() {
				defer wg.Done()
				for n := b.N; n > 0; n-- {
					_ = rand.Uint64()
				}
			}()
		}
		wg.Wait()
	})
}
