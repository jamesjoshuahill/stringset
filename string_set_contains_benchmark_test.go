package stringset_test

import (
	"testing"

	"github.com/jamesjoshuahill/stringset"
	"github.com/pborman/uuid"
)

func BenchmarkSliceForLoopMonkeyAtIndex1(b *testing.B) {
	slice := sliceOfStringsIncludingMonkeys(100, 0)
	for n := 0; n < b.N; n++ {
		for _, element := range slice {
			if element == "monkeys" {
				break
			}
		}
	}
}

func BenchmarkSliceForLoopMonkeyAtIndex25(b *testing.B) {
	slice := sliceOfStringsIncludingMonkeys(100, 24)
	for n := 0; n < b.N; n++ {
		for _, element := range slice {
			if element == "monkeys" {
				break
			}
		}
	}
}

func BenchmarkSliceForLoopMonkeyAtIndex50(b *testing.B) {
	slice := sliceOfStringsIncludingMonkeys(100, 49)
	for n := 0; n < b.N; n++ {
		for _, element := range slice {
			if element == "monkeys" {
				break
			}
		}
	}
}

func BenchmarkSliceForLoopMonkeyAtIndex75(b *testing.B) {
	slice := sliceOfStringsIncludingMonkeys(100, 74)
	for n := 0; n < b.N; n++ {
		for _, element := range slice {
			if element == "monkeys" {
				break
			}
		}
	}
}

func BenchmarkSliceForLoopMonkeyAtIndex100(b *testing.B) {
	slice := sliceOfStringsIncludingMonkeys(100, 99)
	for n := 0; n < b.N; n++ {
		for _, element := range slice {
			if element == "monkeys" {
				break
			}
		}
	}
}

func BenchmarkSliceForLoopNoMonkey(b *testing.B) {
	slice := sliceOfStrings(100)
	for n := 0; n < b.N; n++ {
		for _, element := range slice {
			if element == "monkeys" {
				break
			}
		}
	}
}

func BenchmarkStringSetContainsMonkeyAtIndex1(b *testing.B) {
	set := stringset.New(sliceOfStringsIncludingMonkeys(100, 0)...)
	for n := 0; n < b.N; n++ {
		set.Contains("monkeys")
	}
}

func BenchmarkStringSetContainsMonkeyAtIndex25(b *testing.B) {
	set := stringset.New(sliceOfStringsIncludingMonkeys(100, 24)...)
	for n := 0; n < b.N; n++ {
		set.Contains("monkeys")
	}
}

func BenchmarkStringSetContainsMonkeyAtIndex50(b *testing.B) {
	set := stringset.New(sliceOfStringsIncludingMonkeys(100, 49)...)
	for n := 0; n < b.N; n++ {
		set.Contains("monkeys")
	}
}

func BenchmarkStringSetContainsMonkeyAtIndex75(b *testing.B) {
	set := stringset.New(sliceOfStringsIncludingMonkeys(100, 74)...)
	for n := 0; n < b.N; n++ {
		set.Contains("monkeys")
	}
}

func BenchmarkStringSetContainsMonkeyAtIndex100(b *testing.B) {
	set := stringset.New(sliceOfStringsIncludingMonkeys(100, 99)...)
	for n := 0; n < b.N; n++ {
		set.Contains("monkeys")
	}
}

func BenchmarkStringSetContainsNoMonkey(b *testing.B) {
	set := stringset.New(sliceOfStrings(100)...)
	for n := 0; n < b.N; n++ {
		set.Contains("monkeys")
	}
}

func sliceOfStrings(length int) []string {
	var slice []string

	for i := 0; i < length; i++ {
		slice = append(slice, uuid.New())
	}

	return slice
}

func sliceOfStringsIncludingMonkeys(length, monkeysIndex int) []string {
	var slice []string

	for i := 0; i < length; i++ {
		if i == monkeysIndex {
			slice = append(slice, "monkeys")
		} else {
			slice = append(slice, uuid.New())
		}
	}

	return slice
}
