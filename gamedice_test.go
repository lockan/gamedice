package gamedice

import (
	"math"
	"reflect"
	"testing"
)

func TestRoll(t *testing.T) {
	testroll := 1
	t.Logf("Input is %d", testroll)
	roll := Roll(uint64(testroll))
	t.Logf("Roll is %d", roll)
	var u uint64
	if reflect.TypeOf(u) != reflect.TypeOf(roll) {
		t.Log("Result type is not uint64")
		t.Fail()
	}

	if roll <= 0 {
		t.Log("Result is <= 0")
		t.Fail()
	}

	if roll > math.MaxUint64 {
		t.Log("Result greater than max Uint64")
		t.Fail()
	}
}

func TestZeroInput(t *testing.T) {
	testroll := 0
	t.Logf("Input is %d", testroll)
	roll := Roll(uint64(testroll))
	t.Logf("Roll is %d", roll)
	if 0 != roll {
		t.Fail()
	}
}

func TestMaxUint64(t *testing.T) {
	testroll := 1
	t.Logf("Input is %d", testroll)
	roll := forcedResultRoll(uint64(testroll), math.MaxUint64)
	t.Logf("Roll is %d", roll)
	if roll != math.MaxUint64 {
		t.Fail()
	}
}
