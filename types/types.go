package types

// TODO better package name? (Only required to avoid an import cycle)

// HostCallback represents a function to handle host calls from a waPC guest
type HostCallback func(id uint64, binding string, namespace string, operation string, payload []byte) ([]byte, error)
