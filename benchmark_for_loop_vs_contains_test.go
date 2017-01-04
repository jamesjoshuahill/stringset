package stringset_test

import (
	"testing"

	"github.com/jamesjoshuahill/stringset"
)

var sliceOf100StringsIncludingMonkeysAtIndex1 = sliceOfStringsIncludingMonkeys(100, 0)
var sliceOf100StringsIncludingMonkeysAtIndex25 = sliceOfStringsIncludingMonkeys(100, 24)
var sliceOf100StringsIncludingMonkeysAtIndex50 = sliceOfStringsIncludingMonkeys(100, 49)
var sliceOf100StringsIncludingMonkeysAtIndex75 = sliceOfStringsIncludingMonkeys(100, 74)
var sliceOf100StringsIncludingMonkeysAtIndex100 = sliceOfStringsIncludingMonkeys(100, 99)
var sliceOf100Strings = sliceOfStrings(100)

func BenchmarkSliceForLoopMonkeyAtIndex1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, element := range sliceOf100StringsIncludingMonkeysAtIndex1 {
			if element == "monkeys" {
				break
			}
		}
	}
}

func BenchmarkSliceForLoopMonkeyAtIndex25(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, element := range sliceOf100StringsIncludingMonkeysAtIndex25 {
			if element == "monkeys" {
				break
			}
		}
	}
}

func BenchmarkSliceForLoopMonkeyAtIndex50(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, element := range sliceOf100StringsIncludingMonkeysAtIndex50 {
			if element == "monkeys" {
				break
			}
		}
	}
}

func BenchmarkSliceForLoopMonkeyAtIndex75(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, element := range sliceOf100StringsIncludingMonkeysAtIndex75 {
			if element == "monkeys" {
				break
			}
		}
	}
}

func BenchmarkSliceForLoopMonkeyAtIndex100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, element := range sliceOf100StringsIncludingMonkeysAtIndex100 {
			if element == "monkeys" {
				break
			}
		}
	}
}

func BenchmarkSliceForLoopNoMonkey(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, element := range sliceOf100Strings {
			if element == "monkeys" {
				break
			}
		}
	}
}

func BenchmarkStringSetContainsMonkeyAtIndex1(b *testing.B) {
	set := stringset.New(sliceOf100StringsIncludingMonkeysAtIndex1...)
	for n := 0; n < b.N; n++ {
		set.Contains("monkeys")
	}
}

func BenchmarkStringSetContainsMonkeyAtIndex25(b *testing.B) {
	set := stringset.New(sliceOf100StringsIncludingMonkeysAtIndex25...)
	for n := 0; n < b.N; n++ {
		set.Contains("monkeys")
	}
}

func BenchmarkStringSetContainsMonkeyAtIndex50(b *testing.B) {
	set := stringset.New(sliceOf100StringsIncludingMonkeysAtIndex50...)
	for n := 0; n < b.N; n++ {
		set.Contains("monkeys")
	}
}

func BenchmarkStringSetContainsMonkeyAtIndex75(b *testing.B) {
	set := stringset.New(sliceOf100StringsIncludingMonkeysAtIndex75...)
	for n := 0; n < b.N; n++ {
		set.Contains("monkeys")
	}
}

func BenchmarkStringSetContainsMonkeyAtIndex100(b *testing.B) {
	set := stringset.New(sliceOf100StringsIncludingMonkeysAtIndex100...)
	for n := 0; n < b.N; n++ {
		set.Contains("monkeys")
	}
}

func BenchmarkStringSetContainsNoMonkey(b *testing.B) {
	set := stringset.New(sliceOf100Strings...)
	for n := 0; n < b.N; n++ {
		set.Contains("monkeys")
	}
}
