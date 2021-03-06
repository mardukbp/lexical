package parse

import (
	"io"
	"testing"

	"github.com/a-h/lexical/input"
)

func TestAnyRune(t *testing.T) {
	tests := []struct {
		input         string
		expected      bool
		expectedError error
	}{
		{
			input:    "A",
			expected: true,
		},
		{
			input:         "",
			expected:      false,
			expectedError: io.EOF,
		},
	}

	for i, test := range tests {
		pi := input.NewFromString(test.input)
		parser := AnyRune()
		result := parser(pi)
		actual := result.Success
		if actual != test.expected {
			t.Errorf("test %v: for input '%v' expected %v but got %v", i, test.input, test.expected, actual)
		}
		if result.Error != test.expectedError {
			t.Errorf("test %v: for input '%v' expected error '%v' but got '%v'", i, test.input, test.expectedError, result.Error)
		}
	}
}

func BenchmarkAnyRune(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		parser := AnyRune()
		parser(input.NewFromString("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
	}
}
