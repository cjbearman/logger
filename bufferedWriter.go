package logger

import (
	"errors"
	"io"
	"sync"
)

var errClosed = errors.New("closed")

// bufferWriter
type BufferedWriter struct {
	ch     chan []byte
	clch   chan interface{}
	target io.Writer
	closed bool
	lock   sync.Mutex
}

// Write data to a the buffer
func (w *BufferedWriter) Write(data []byte) (n int, err error) {
	w.lock.Lock()
	defer w.lock.Unlock()
	if w.closed {
		return 0, errClosed
	}
	w.ch <- data
	return len(data), nil
}

// Close should be called when you are done with this writer
func (w *BufferedWriter) Close() error {
	w.lock.Lock()
	defer w.lock.Unlock()
	if w.closed {
		return nil
	}
	close(w.ch)
	w.closed = true
	<-w.clch
	return nil
}

// run implements the go routine that does the terminal logging
func (w *BufferedWriter) run() {
	for data := range w.ch {
		w.target.Write(data)
	}
	close(w.clch)
}

// NewBufferedWriter returns a writer that will avoid blocking of calling threads
// by delegating the logging to a sub-goroutine
// The danger of this is loosing logs if your program panics
// so make sure and call Close reliably to ensure all logs are output
func NewBufferedWriter(target io.Writer, bufSize int) *BufferedWriter {
	bw := &BufferedWriter{target: target, ch: make(chan []byte, bufSize), clch: make(chan interface{})}
	go bw.run()
	return bw
}
