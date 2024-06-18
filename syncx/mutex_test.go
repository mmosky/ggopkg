package syncx

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
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

func HammerMutex(m *Mutex, loops int, cdone chan bool) {
	for i := 0; i < loops; i++ {
		if i%3 == 0 {
			if m.TryLock() {
				m.Unlock()
			}
			continue
		}
		m.Lock()
		m.Unlock()
	}
	cdone <- true
}

func TestMutexStd(t *testing.T) {
	if n := runtime.SetMutexProfileFraction(1); n != 0 {
		t.Logf("got mutexrate %d expected 0", n)
	}
	defer runtime.SetMutexProfileFraction(0)

	m := new(Mutex)

	m.Lock()
	if m.TryLock() {
		t.Fatalf("TryLock succeeded with mutex locked")
	}
	m.Unlock()
	if !m.TryLock() {
		t.Fatalf("TryLock failed with mutex unlocked")
	}
	m.Unlock()

	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go HammerMutex(m, 1000, c)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

var misuseTests = []struct {
	name string
	f    func()
}{
	{
		"Mutex.Unlock",
		func() {
			var mu Mutex
			mu.Unlock()
		},
	},
	{
		"Mutex.Unlock2",
		func() {
			var mu Mutex
			mu.Lock()
			mu.Unlock()
			mu.Unlock()
		},
	},
	// {
	// 	"RWMutex.Unlock",
	// 	func() {
	// 		var mu RWMutex
	// 		mu.Unlock()
	// 	},
	// },
	// {
	// 	"RWMutex.Unlock2",
	// 	func() {
	// 		var mu RWMutex
	// 		mu.RLock()
	// 		mu.Unlock()
	// 	},
	// },
	// {
	// 	"RWMutex.Unlock3",
	// 	func() {
	// 		var mu RWMutex
	// 		mu.Lock()
	// 		mu.Unlock()
	// 		mu.Unlock()
	// 	},
	// },
	// {
	// 	"RWMutex.RUnlock",
	// 	func() {
	// 		var mu RWMutex
	// 		mu.RUnlock()
	// 	},
	// },
	// {
	// 	"RWMutex.RUnlock2",
	// 	func() {
	// 		var mu RWMutex
	// 		mu.Lock()
	// 		mu.RUnlock()
	// 	},
	// },
	// {
	// 	"RWMutex.RUnlock3",
	// 	func() {
	// 		var mu RWMutex
	// 		mu.RLock()
	// 		mu.RUnlock()
	// 		mu.RUnlock()
	// 	},
	// },
}

func init() {
	if len(os.Args) == 3 && os.Args[1] == "TESTMISUSE" {
		for _, test := range misuseTests {
			if test.name == os.Args[2] {
				func() {
					defer func() { recover() }()
					test.f()
				}()
				fmt.Printf("test completed\n")
				os.Exit(0)
			}
		}
		fmt.Printf("unknown test\n")
		os.Exit(0)
	}
	SetGlobalTimeout(time.Second)
}

// MustHaveExec checks that the current system can start new processes
// using os.StartProcess or (more commonly) exec.Command.
// If not, MustHaveExec calls t.Skip with an explanation.
//
// On some platforms MustHaveExec checks for exec support by re-executing the
// current executable, which must be a binary built by 'go test'.
// We intentionally do not provide a HasExec function because of the risk of
// inappropriate recursion in TestMain functions.
//
// To check for exec support outside of a test, just try to exec the command.
// If exec is not supported, testenv.SyscallIsNotSupported will return true
// for the resulting error.
func MustHaveExec(t testing.TB) {
	tryExecOnce.Do(func() {
		tryExecErr = tryExec()
	})
	if tryExecErr != nil {
		t.Skipf("skipping test: cannot exec subprocess on %s/%s: %v", runtime.GOOS, runtime.GOARCH, tryExecErr)
	}
}

var (
	tryExecOnce sync.Once
	tryExecErr  error
	origEnv     = os.Environ()
)

func tryExec() error {
	switch runtime.GOOS {
	case "wasip1", "js", "ios":
	default:
		// Assume that exec always works on non-mobile platforms and Android.
		return nil
	}

	// ios has an exec syscall but on real iOS devices it might return a
	// permission error. In an emulated environment (such as a Corellium host)
	// it might succeed, so if we need to exec we'll just have to try it and
	// find out.
	//
	// As of 2023-04-19 wasip1 and js don't have exec syscalls at all, but we
	// may as well use the same path so that this branch can be tested without
	// an ios environment.

	if !testing.Testing() {
		// This isn't a standard 'go test' binary, so we don't know how to
		// self-exec in a way that should succeed without side effects.
		// Just forget it.
		return errors.New("can't probe for exec support with a non-test executable")
	}

	// We know that this is a test executable. We should be able to run it with a
	// no-op flag to check for overall exec support.
	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("can't probe for exec support: %w", err)
	}
	cmd := exec.Command(exe, "-test.list=^$")
	cmd.Env = origEnv
	return cmd.Run()
}

func TestMutexMisuse(t *testing.T) {
	MustHaveExec(t)
	for _, test := range misuseTests {
		out, err := exec.Command(os.Args[0], "TESTMISUSE", test.name).CombinedOutput()
		if err == nil || !strings.Contains(string(out), "unlocked") {
			t.Errorf("%s: did not find failure with message about unlocked lock: %s\n%s\n", test.name, err, out)
		}
	}
}

func TestMutexFairness(t *testing.T) {
	var mu Mutex
	stop := make(chan bool)
	defer close(stop)
	go func() {
		for {
			mu.Lock()
			time.Sleep(100 * time.Microsecond)
			mu.Unlock()
			select {
			case <-stop:
				return
			default:
			}
		}
	}()
	done := make(chan bool, 1)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Microsecond)
			mu.Lock()
			mu.Unlock()
		}
		done <- true
	}()
	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatalf("can't acquire Mutex in 10 seconds")
	}
}

func BenchmarkMutexUncontended(b *testing.B) {
	type PaddedMutex struct {
		Mutex
		pad [128]uint8
	}
	b.RunParallel(func(pb *testing.PB) {
		var mu PaddedMutex
		for pb.Next() {
			mu.Lock()
			mu.Unlock()
		}
	})
}

func benchmarkMutex(b *testing.B, slack, work bool) {
	var mu Mutex
	if slack {
		b.SetParallelism(10)
	}
	b.RunParallel(func(pb *testing.PB) {
		foo := 0
		for pb.Next() {
			mu.Lock()
			mu.Unlock()
			if work {
				for i := 0; i < 100; i++ {
					foo *= 2
					foo /= 2
				}
			}
		}
		_ = foo
	})
}

func BenchmarkMutexStd(b *testing.B) {
	benchmarkMutex(b, false, false)
}

func BenchmarkMutexSlack(b *testing.B) {
	benchmarkMutex(b, true, false)
}

func BenchmarkMutexWork(b *testing.B) {
	benchmarkMutex(b, false, true)
}

func BenchmarkMutexWorkSlack(b *testing.B) {
	benchmarkMutex(b, true, true)
}

func BenchmarkMutexNoSpin(b *testing.B) {
	// This benchmark models a situation where spinning in the mutex should be
	// non-profitable and allows to confirm that spinning does not do harm.
	// To achieve this we create excess of goroutines most of which do local work.
	// These goroutines yield during local work, so that switching from
	// a blocked goroutine to other goroutines is profitable.
	// As a matter of fact, this benchmark still triggers some spinning in the mutex.
	var m Mutex
	var acc0, acc1 uint64
	b.SetParallelism(4)
	b.RunParallel(func(pb *testing.PB) {
		c := make(chan bool)
		var data [4 << 10]uint64
		for i := 0; pb.Next(); i++ {
			if i%4 == 0 {
				m.Lock()
				acc0 -= 100
				acc1 += 100
				m.Unlock()
			} else {
				for i := 0; i < len(data); i += 4 {
					data[i]++
				}
				// Elaborate way to say runtime.Gosched
				// that does not put the goroutine onto global runq.
				go func() {
					c <- true
				}()
				<-c
			}
		}
	})
}

func BenchmarkMutexSpin(b *testing.B) {
	// This benchmark models a situation where spinning in the mutex should be
	// profitable. To achieve this we create a goroutine per-proc.
	// These goroutines access considerable amount of local data so that
	// unnecessary rescheduling is penalized by cache misses.
	var m Mutex
	var acc0, acc1 uint64
	b.RunParallel(func(pb *testing.PB) {
		var data [16 << 10]uint64
		for i := 0; pb.Next(); i++ {
			m.Lock()
			acc0 -= 100
			acc1 += 100
			m.Unlock()
			for i := 0; i < len(data); i += 4 {
				data[i]++
			}
		}
	})
}
