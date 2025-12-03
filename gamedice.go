package gamedice

import (
	"fmt"
	"math"
	rand "math/rand/v2"
	"strings"
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

func sanitize(inputString string) string {
	clean := strings.TrimSpace(inputString)
	clean = strings.Join(strings.Fields(clean), "")
	clean = strings.ToLower(clean)
	return clean
}

// Parse Roll String in format XdN +/- Y
func ParseRoll(rollstring string) string {
	//TODO - implement me
	// Considerations: should handle both upper and lower case by converting to lowercase.
	// counts and sides cannot be zero (modifiers can be even if it's silly)
	// Handle whitespace in input by stripping it out.
	// Handle different ordered tokens e.g. mdn+x, x+mdn, mdn+kdn

	// Variations:
	// dN - both upper and lower case
	// MdN
	// dN+/-X
	// MdN+/-X
	// MdN+KdN
	// X+/-MdN

	// Algo:
	// Given input string s
	// Trim all whitespace from s and convert to lowercase
	// Init tokens []string
	// walk s from index 0 ... to len(s)
	// Each time a separator character is found:
	// 		add the preceding token to tokens []string,
	// 		followed by the separator as a separate token
	// stop at end of of s
	// Walk tokens array. For each token:
	// 		If token is not a separator, check for 'd' in string. If d is found:
	// 			if d is the first character RollN and return result
	//			if d is NOT the first character RollBatch([0][1]) and return result
	//		If token is a separator, determine which math operation, move to next token, and sum with previous result
	// Finally: return total
	// var separators = [...]rune{'+', '-'}
	// var d rune = 'd'
	// var tokens = [...]string{}
	rollstring = sanitize(rollstring)
	return rollstring
}

// ===== misc private test functions
func forcedResultRoll(sides uint64, forced uint64) uint64 {
	result := Roll(sides)
	result = forced // override to the value we want.
	return result
}
