package gamedice

import (
	"fmt"
	"math"
	rand "math/rand/v2"
)

// Bad Input Handler (e.g. if 0 is provided for a uint64N call.)
// TODO: proper implementation, update tests
func badInputHandler(i any) uint64 {
	rec := recover()
	var result uint64
	if rec != nil {
		fmt.Printf("Handled bad input value: %d", i)
		result = 0
	} else {
		result = 1
	}
	return result
}

// Roll dN die
func Roll(sides uint64) uint64 {
	defer badInputHandler(sides)
	result := rand.Uint64N(sides)

	if result >= math.MaxUint64 {
		return math.MaxUint64
	} else {
		return result + 1
	}
}

// Roll MdN
func RollN(count uint64, sides uint64) uint64 {
	defer badInputHandler(count)
	defer badInputHandler(sides)

	var sumroll uint64
	values := RollBatch(count, sides)
	for i := range values {
		sumroll = sumroll + values[i]
	}
	return sumroll
}

// Roll MdN as batch of unique dice values
func RollBatch(count uint64, sides uint64) []uint64 {
	defer badInputHandler(count)
	defer badInputHandler(sides)

	batch := make([]uint64, count)
	for i := 0; uint64(i) < count; i++ {
		batch[i] = Roll(sides)
	}
	return batch
}

//Parse Roll String in format XdN +/- Y

// ===== misc private test functions
func forcedResultRoll(sides uint64, forced uint64) uint64 {
	result := Roll(sides)
	result = forced // override to the value we want.
	return result
}
