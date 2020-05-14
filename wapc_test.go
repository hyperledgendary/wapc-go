package wapc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This test is based on the readme example
func TestNewHostExample(t *testing.T) {
	moduleBytes := []byte("mock wasm bytes!")

	// TODO mock Life VM somehow?
	host := NewHost(func(id uint64, bd string, ns string, op string, payload []byte) ([]byte, error) {
		fmt.Printf("Guest %d invoked %s %s %s with payload of %d bytes\n", id, bd, ns, op, int64(len(payload)))
		return nil, nil
	}, moduleBytes)
	assert.NotNil(t, host)

	// TODO implement call method!
	// result := host.call("wapc:sample!Hello", []byte("this is a test"))
	result := "hello world!"
	if string(result) != "hello world!" {
		panic("Unexpected response from guest")
	}
}
