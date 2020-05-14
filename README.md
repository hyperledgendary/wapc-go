# waPC Go

This is the Golang implementation of the [waPC](https://github.com/wapc) standard for WebAssembly host runtimes. It allows any WebAssembly module to be loaded as a guest and receive requests for invocation as well as to make its own function requests of the host. To begin with this implementation is only intended to work with the [Perlin Life WebAssembly VM](https://github.com/perlin-network/life)

## Example

```
func main() {
    moduleBytes, err := loadFile()
    if err != nil {
        panic(err)
    }

    host := wapc.NewHost(func(id uint64, bd string, ns string, op string, payload []byte) ([]byte, error) {
		fmt.Printf("Guest %d invoked %s %s %s with payload of %d bytes\n", id, bd, ns, op, int64(len(payload)))
		return nil, nil
	}, moduleBytes)

	result := host.call("wapc:sample!Hello", []byte("this is a test"))
    if string(result) != "hello world!" {
        panic("Unexpected response from guest")
    }

	return
}
```

## Notes

waPC is _reactive_.
Guest modules cannot initiate host calls without first handling a call initiated by the host.
waPC will not automatically invoke any start functions--that decision is up to the waPC library consumer.
Guest modules can synchronously make as many host calls as they like, but keep in mind that if a host call takes too long or fails, it'll cause the original guest call to also fail.

In summary, keep `host_callback` functions fast and resilient, and do not spawn new threads within `host_callback` unless you must (and can synchronize memory access) because waPC assumes a single-threaded execution environment.
The `host_callback` function intentionally has no references to the WebAssembly module bytes or the running instance.
