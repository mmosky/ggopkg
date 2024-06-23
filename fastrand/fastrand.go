package fastrand

import (
	_ "unsafe"
)

// Uint64 returns a random uint64 from the per-m chacha8 state.
// It is safe for concurrent use by multiple goroutines.
// Comparing to math/rand.Uint64, fastrand.Uint64 can save
// about 20% time.
//
//go:linkname Uint64 runtime.rand
func Uint64() uint64
