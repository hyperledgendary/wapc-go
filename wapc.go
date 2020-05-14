package wapc

import (
	"github.com/hyperledgendary/wapc-go/internal"
	"github.com/hyperledgendary/wapc-go/types"
)

// Host represents a WebAssembly host runtime for waPC-compliant WebAssembly modules
//
// Use an instance of this struct to provide a means of invoking procedure calls by
// specifying an operation name and a set of bytes representing the opaque operation payload.
// `Host` makes no assumptions about the contents or format of either the payload or the
// operation name.
type Host struct {
	resolver *internal.WapcResolver
}

// NewHost nice new host you have there
func NewHost(cb types.HostCallback, buf []byte) *Host {
	resolver := internal.NewWapcResolver(cb)

	host := new(Host)
	host.resolver = resolver

	// TODO create Life Wasm VM

	return host
}

// TODO implement call method
