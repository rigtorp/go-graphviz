package graphviz

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

//go:embed testdata/graph.dot
var testGraph string

func TestGraphviz(t *testing.T) {
	entries, err := os.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if filepath.Ext(entry.Name()) != ".dot" {
			continue
		}
		t.Run(entry.Name(), func(t *testing.T) {
			in, err := os.Open(filepath.Join("testdata", entry.Name()))
			if err != nil {
				t.Fatal(err)
			}
			want, err := os.ReadFile(filepath.Join("testdata", strings.TrimSuffix(entry.Name(), ".dot")+".svg"))
			if err != nil {
				t.Fatal(err)
			}
			got := bytes.Buffer{}

			err = Render(in, &got)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(got.Bytes(), want) {
				t.Fatalf("%s:\n\nwant:\n%s\n\ngot:\n%s\n", entry.Name(), want, got.String())
			}
		})
	}
}

func BenchmarkGraphviz(b *testing.B) {
	b.RunParallel(benchStep)
}

func benchStep(pb *testing.PB) {
	for pb.Next() {
		result, err := RenderString(testGraph)
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(io.Discard, result)
	}
}
