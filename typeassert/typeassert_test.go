package typeassert

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func Test_fooToBar(t *testing.T) {
	f := Foo{s: "The Foo string"}
	b := fooToBar(f)
	spew.Dump(f, b)
}
