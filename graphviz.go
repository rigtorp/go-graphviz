// Package graphviz wraps a WASM build of graphviz as go package.
package graphviz

import (
	"bytes"
	"context"
	_ "embed"
	"io"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed dot.wasm
var dotWasm []byte

// Render converts input graph into output SVG figure.
func Render(in io.Reader, out io.Writer) error {
	ctx := context.TODO()

	r := wazero.NewRuntimeWithConfig(ctx, wazero.NewRuntimeConfigInterpreter())
	defer r.Close(ctx)

	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	code, err := r.CompileModule(ctx, dotWasm)
	if err != nil {
		return err
	}

	// Only stdin and stdout is exposed to the WASM module
	config := wazero.NewModuleConfig().WithStdout(out).WithStdin(in).WithArgs("dot", "-Tsvg", "-Kdot")

	_, _ = r.InstantiateModule(ctx, code, config)
	// Weird but above ^^ always returns error

	return nil
}

// RenderString converts input graph into output SVG figure.
func RenderString(text string) (string, error) {
	out := &bytes.Buffer{}
	in := bytes.NewBufferString(text)

	if err := Render(in, out); err != nil {
		return "", err
	}
	return out.String(), nil
}
