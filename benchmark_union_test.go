package stringset_test

import (
	"testing"

	"github.com/jamesjoshuahill/stringset"
)

func BenchmarkStringSetUnion(b *testing.B) {
	set := stringset.New(sliceOfStrings(100)...)
	other := stringset.New(sliceOfStrings(100)...)
	for n := 0; n < b.N; n++ {
		set.Union(other)
	}
}
