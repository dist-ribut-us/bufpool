// Package bufpool provides a pool of bytes.Buffer.
package bufpool

import (
	"bytes"
	"sync"
)

var compressionBuffer = &sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(nil)
	},
}

// Put a Buffer in the pool
func Put(b *bytes.Buffer) {
	b.Reset()
	compressionBuffer.Put(b)
}

// Get a Buffer from the pool
func Get() *bytes.Buffer {
	return compressionBuffer.Get().(*bytes.Buffer)
}
