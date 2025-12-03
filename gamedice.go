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

func sanitize(input string) string {
	// strip all whitespace, convert to lowercase.
	clean := strings.TrimSpace(input)
	clean = strings.Join(strings.Fields(clean), "")
	clean = strings.ToLower(clean)
	return clean
}

func isValidRoll(rollstring string) bool {
	// check for only allowable characters in a sanitized roll string
	var validchars string = "0123456789+-d"
	// check for invalid characters
	for i := range rollstring {
		if !strings.Contains(validchars, string(rollstring[i])) {
			return false
		}
	}
	// ensure we have at least one instance of a 'd'
	if !strings.ContainsRune(rollstring, 'd') {
		return false
	}
	return true
}

func tokenize(rollstring string) []string {
	// Algo:
	// Given input string s
	// Init tokens []string
	// walk s from index 0 ... to len(s)
	// Each time a separator character is found:
	// 		add the preceding token to tokens []string,
	// 		followed by the separator as a separate token
	// stop at end of of s, return tokenized slice
	tokens := []string{}
	return tokens
}

// Parse Roll String in format XdN +/- Y
func ParseRoll(rollstring string) string {
	//TODO - implement me
	// counts and sides cannot be zero (modifiers can be even if it's silly)
	// Handle different ordered tokens e.g. mdn+x, x+mdn, mdn+kdn
	// rollstring could result in negative sums; If it does return zero.

	// Variations:
	// dN - both upper and lower case
	// MdN
	// dN+/-X
	// MdN+/-X
	// MdN+KdN
	// X+/-MdN

	// Algo:
	// Given a slice of token string
	// For each token:
	// 		If token is not a separator, check for 'd' in string. If d is found:
	// 			if d is the first character RollN and return result
	//			if d is NOT the first character RollBatch([0][1]) and return result
	//		If token is a separator, determine which math operation, move to next token, and sum with previous result
	// Finally: return total
	// var separators = [...]rune{'+', '-'}
	// var d rune = 'd'
	// var tokens = [...]string{}
	rollstring = sanitize(rollstring)
	//tokens := tokenize(rollstring)

	return rollstring
}

// ===== misc private test functions
func forcedResultRoll(sides uint64, forced uint64) uint64 {
	result := Roll(sides)
	result = forced // override to the value we want.
	return result
}
