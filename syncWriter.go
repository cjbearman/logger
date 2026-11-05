package logger

import (
	"io"
	"sync"
)

// syncWriter can be used to synchronize logs to prevent
// inter-mingling from multiple threads with zero possibility of log loss
// but with logs written from the writing thread, and thus less performant
// than bufferedWriter
type syncWriter struct {
	lock   sync.Mutex
	target io.Writer
}

func (w *syncWriter) Write(data []byte) (n int, err error) {
	w.lock.Lock()
	defer w.lock.Unlock()
	return w.target.Write(data)
}

// NewSyncWriter returns a writer that ensures that logs are output
// sequentially, and do not intermingle on stdout
// Use if logging from multiple goroutines
func NewSyncWriter(target io.Writer) io.Writer {
	return &syncWriter{target: target}
}
