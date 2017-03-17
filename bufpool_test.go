package bufpool

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBufPool(t *testing.T) {
	b := Get()
	assert.NotNil(t, b)
	Put(b)
}
