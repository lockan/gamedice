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

func TestRollBatch(t *testing.T) {
	var count, sides uint64
	count = 4
	sides = 6
	t.Logf("Roll is %dD%d", count, sides)
	batch := RollBatch(count, sides)
	t.Log("Batch Result is:")
	t.Log(batch)

	if uint64(len(batch)) != count {
		t.Log("Got incorrect batch size as result!")
		t.Fail()
	}

	for i := range batch {
		if batch[i] <= 0 || batch[i] > math.MaxUint64 {
			t.Logf("One of the results is an invalid value: %d", batch[i])
			t.Fail()
		}
	}
}

func TestRollN(t *testing.T) {
	var count, sides uint64
	count = 4
	sides = 6
	t.Logf("Roll is %dD%d", count, sides)
	roll := RollN(uint64(count), sides)
	t.Logf("Result is %d", roll)

	if roll <= 0 || roll > math.MaxUint64 {
		t.Log("Rolled an invalid value, somehow!!")
		t.Fail()
	}

	if roll > count*sides {
		t.Log("Rolled value is greater than maximum possible roll.")
		t.Fail()
	}
}

func TestSanitize(t *testing.T) {
	// Test string sanitization
	var teststring string = "    4 D 6 + 	2			"
	var expected string = "4d6+2"
	var result string = sanitize(teststring)
	t.Logf("Input string: %s; Expecting: %s", teststring, expected)
	t.Logf("Parsed string: %s", result)
	if result != expected {
		t.Fail()
	}
}

func TestValidate(t *testing.T) {
	var valid = "10-2d10"
	var invalid = "40xd7+5yyy"
	var hasD = "d4"
	var missingD = "1+2"
	if !isValidRoll(valid) {
		t.Logf("isValidRoll returned a false negative!")
		t.Fail()
	}
	if isValidRoll(invalid) {
		t.Logf("isValidRoll got a false positive!")
		t.Fail()
	}
	if !isValidRoll(hasD) {
		t.Logf("isValidRoll: could not find character 'd'")
		t.Fail()
	}
	if isValidRoll(missingD) {
		t.Logf("isValidRoll: found character 'd' in string that has")
		t.Fail()
	}
}
