# Graphviz for Go

This package wraps [Graphviz](https://graphviz.org/) as a Go package. It doesn't
rely on cgo or an external binary, instead using the
[wazero](https://wazero.io/) WebAssembly runtime to embed Graphviz.

Wazero is configured to sandbox Graphviz such that Graphviz only has access to
the input and output data.

## Performance

For a small Graphviz graph, go-graphviz takes about 51ms:

```shell
$ go test -test.bench .
goos: linux
goarch: amd64
pkg: github.com/rigtorp/go-graphviz
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkGraphviz-8           24          51285746 ns/op
PASS
ok      github.com/rigtorp/go-graphviz  3.057s
```

## Building Graphviz for WASM

Download [wasi-sdk](https://github.com/WebAssembly/wasi-sdk) and extract it
somewhere.

Next build Graphviz:

```shell
mkdir build
cd build
cmake .. -DWASI_SDK_PREFIX=/path/to/wasi/sdk
ninja
```

## Acknowledgements

I got the idea for this approach of embedding an external program in Go using
WebAssembly from the blog post [The carcinization of Go
programs](https://xeiaso.net/blog/carcinization-golang).
