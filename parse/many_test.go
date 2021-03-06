package parse

import (
	"testing"

	"github.com/a-h/lexical/input"
)

func TestMany(t *testing.T) {
	tests := []struct {
		input         string
		parser        Function
		expectedMatch bool
		expectedItem  interface{}
	}{
		{
			input:         "AAAAAAAA",
			parser:        Many(WithStringConcatCombiner, 0, 500, Rune('A')),
			expectedMatch: true,
			expectedItem:  "AAAAAAAA",
		},
		{
			input:         "AAABBB",
			parser:        Many(WithStringConcatCombiner, 0, 500, Rune('A')),
			expectedMatch: true,
			expectedItem:  "AAA",
		},
		{
			input:         "AAABBB",
			parser:        Many(WithStringConcatCombiner, 0, 500, Or(Rune('A'), Rune('B'))),
			expectedMatch: true,
			expectedItem:  "AAABBB",
		},
		{
			input:         "AAABBB",
			parser:        Many(WithStringConcatCombiner, 4, 500, Rune('A')),
			expectedMatch: false,
		},
		{
			input:         "AAABBB",
			parser:        Many(WithStringConcatCombiner, 1, 2, Rune('A')),
			expectedMatch: true,
			expectedItem:  "AA",
		},
		{
			input:         "1",
			parser:        Many(WithIntegerCombiner, 1, 2, ZeroToNine),
			expectedMatch: true,
			expectedItem:  1,
		},
		{
			input:         "12",
			parser:        Many(WithIntegerCombiner, 1, 2, ZeroToNine),
			expectedMatch: true,
			expectedItem:  12,
		},
		{
			input:         "12",
			parser:        AtLeast(WithIntegerCombiner, 1, ZeroToNine),
			expectedMatch: true,
			expectedItem:  12,
		},
		{
			input:         "12",
			parser:        AtMost(WithIntegerCombiner, 1, ZeroToNine),
			expectedMatch: true,
			expectedItem:  1,
		},
		{
			input:         "12",
			parser:        Times(WithIntegerCombiner, 1, ZeroToNine),
			expectedMatch: true,
			expectedItem:  1,
		},
	}

	for i, test := range tests {
		pi := input.NewFromString(test.input)
		result := test.parser(pi)
		actualMatch := result.Success
		if actualMatch != test.expectedMatch {
			t.Errorf("test %v: for input '%v' expected %v but got %v", i, test.input, test.expectedMatch, actualMatch)
		}
		if test.expectedMatch && result.Item != test.expectedItem {
			t.Errorf("test %v: for input '%v' expected to capture '%v' but got '%v'", i, test.input, test.expectedItem, result.Item)
		}
	}
}

func BenchmarkMany(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		parser := Many(WithStringConcatCombiner, 1, 1, Rune('A'))
		parser(input.NewFromString("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
	}
}
