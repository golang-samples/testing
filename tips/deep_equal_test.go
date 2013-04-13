package example_test

import (
	"reflect"
	"testing"
)

type MyType struct {
	foo int
	bar float32
}

func TestDeepEqual(t *testing.T) {
	a := MyType{1, 3.5}
	b := MyType{1, 3.5}

	if !reflect.DeepEqual(a, b) {
		t.Errorf("mismatch. got: %v, %v", a, b)
	}
}
