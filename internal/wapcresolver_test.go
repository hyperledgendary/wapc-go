package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHostCallback(t *testing.T) {
	resolver := NewWapcResolver(func(id uint64, bd string, ns string, op string, payload []byte) ([]byte, error) {
		result := fmt.Sprintf("Guest %d invoked %s %s %s with payload of %d bytes", id, bd, ns, op, int64(len(payload)))

		return []byte(result), nil
	})

	cbResult, err := resolver.hostCallback(42, "bd", "ns", "op", []byte("TESTING"))
	assert.Nil(t, err)
	assert.Equal(t, string(cbResult), "Guest 42 invoked bd ns op with payload of 7 bytes")
}

// TODO test functions called by guest, exported by host:
// __console_log
// __host_call
// __guest_request
// __host_response
// __host_response_len
// __guest_response
// __guest_error
// __host_error
// __host_error_len
