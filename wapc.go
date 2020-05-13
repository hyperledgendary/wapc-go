package wapc

// Host A WebAssembly host runtime for waPC-compliant WebAssembly modules
//
// Use an instance of this struct to provide a means of invoking procedure calls by
// specifying an operation name and a set of bytes representing the opaque operation payload.
// `Host` makes no assumptions about the contents or format of either the payload or the
// operation name.
type Host struct {
}
