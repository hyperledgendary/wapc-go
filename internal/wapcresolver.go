package internal

import (
	"github.com/hyperledgendary/wapc-go/types"
)

// WapcResolver defines the waPC imports for WebAssembly.
type WapcResolver struct {
	hostCallback types.HostCallback
}

// NewWapcResolver creates a new resolver with the supplied host callback function
func NewWapcResolver(cb types.HostCallback) *WapcResolver {
	resolver := new(WapcResolver)
	resolver.hostCallback = cb

	return resolver
}

// TODO implement required methods to resolve waPC imports
