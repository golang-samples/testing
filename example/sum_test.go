package example_test

import (
	. "github.com/golang-sample/testing/example"
	"testing"
)

type sumTest struct {
	in1, in2, out int
}

// This is data provider
var sumTests = []sumTest{
	sumTest{1, 2, 3},
	sumTest{5, 2, 7},
}

func TestSum(t *testing.T) {
	for _, test := range sumTests {
		out := Sum(test.in1, test.in2)
		if out != test.out {
			t.Errorf("Sum(%v, %v) = %v, want %v", test.in1, test.in2, out, test.out)
		}
	}
}

// You can run this benchmark using "go test -test.bench Sum"
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 3)
	}
}
