package syncx

// GetGoID returns the id of current goroutine.
// func GetGoID() uint64 {
// 	b := make([]byte, 64)
// 	b = b[:runtime.Stack(b, false)]
// 	b = bytes.TrimPrefix(b, []byte("goroutine "))
// 	b = b[:bytes.IndexByte(b, ' ')]
// 	n, _ := strconv.ParseUint(string(b), 10, 64)
// 	return n
// }

// type goIDTyp struct {}

// var goIDIncr uint64

// func WithGoID(ctx context.Context) context.Context {
// 	id := atomic.AddUint64(&goIDIncr, 1)
// 	return context.WithValue(ctx, goIDTyp{}, id)
// }
