package gamedice

import (
	"fmt"
	"math"
	rand "math/rand/v2"
	"strconv"
	"strings"
)

const VALIDCHARS string = "0123456789+-d"
const OPERATORS string = "+-"

// Zero Input Handler (e.g. if 0 is provided for a uint64N call.)
func zeroInputHandler(i uint64) uint64 {
	rec := recover()
	if rec != nil {
		fmt.Printf("Handled bad input value: %d", i)
		return 0
	}
	return i
}

// Test for overflow. uint64 can't be a negative, so overflows to zero.
// If a summed roll results in a value less than the previous we can assume overflow has occurred.
// We're summing game dice here, so realistically if a number is that large: just return MaxUint64.
func overflowHandler(current uint64, previous uint64) uint64 {
	// check for overlow
	if current < previous {
		fmt.Printf("uint64 overflow detected! returning max uint64.MaxValue")
		return math.MaxUint64
	} else {
		return current
	}
}

// Roll dN die
func Roll(sides uint64) uint64 {
	defer zeroInputHandler(sides)
	result := rand.Uint64N(sides)
	return result + 1
}

// Roll MdN dice and return the total value
// Returns a single summed result
func RollN(count uint64, sides uint64) uint64 {
	defer zeroInputHandler(count)
	defer zeroInputHandler(sides)

	var sumroll uint64
	values := RollBatch(count, sides)
	for i := range values {
		nextsum := sumroll + values[i]
		defer overflowHandler(nextsum, sumroll)
		sumroll = nextsum
	}
	return sumroll
}

// Roll MdN as batch of M unique dice values of type dN
// Returns multiple independent results, rather than a single sum
func RollBatch(count uint64, sides uint64) []uint64 {
	defer zeroInputHandler(count)
	defer zeroInputHandler(sides)

	batch := make([]uint64, count)
	for i := 0; uint64(i) < count; i++ {
		batch[i] = Roll(sides)
	}
	return batch
}

// Carries out a complex series of rolls from a list of roll tokens. e.g. parseRoll() results.
// Returns the total result.
func RollComplex(tokens []string) uint64 {
	// TODO: Implement
	values := []uint64{}
	for i := range tokens {
		token := tokens[i]
		fmt.Printf("%s", token)
		if isDieRoll(token) {
			// TODO: IMPLEMENT
			count, sides, err := ParseRoll(token)
			if err != nil {
				defer panic(err)
			}
			values = append(values, RollN(count, sides))
		} else if isOperator(token) {
			// TODO: IMPLEMENT
			// 1. get next token
			// 2. ... profit?
			values = append(values, RollN(1, 11))
		} else {
			// TODO: Handle static values
			// ModifyRoll goes here ... I think?
			values = append(values, 13)
		}
	}
	var result uint64
	return result
}

// Adds or subtracts a static value from a previous subtotal
func ModifyRoll(operator rune, subtotal uint64, staticvalue uint64) uint64 {
	var result uint64
	switch operator {
	case '+':
		result = subtotal + staticvalue
		defer overflowHandler(result, subtotal)
	case '-':
		result = subtotal - staticvalue
		// TODO: do I also need an underflow handler?
	default:
		defer panic(fmt.Sprintf("Not a valid operator: %s", string(operator)))
	}
	return result
}

func sanitize(input string) string {
	// strip all whitespace, convert to lowercase.
	clean := strings.TrimSpace(input)               // removes leading/trailing space
	clean = strings.Join(strings.Fields(clean), "") // removes internal white space
	clean = strings.ToLower(clean)
	return clean
}

func isDieRoll(rollstring string) bool {
	// ensure we have at least one instance of a 'd'
	if !strings.ContainsRune(rollstring, 'd') {
		return false
	}
	return true
}

func isValidRoll(rollstring string) bool {
	// check for only allowable characters in a sanitized roll string
	// check for invalid characters
	for i := range rollstring {
		if !strings.Contains(VALIDCHARS, string(rollstring[i])) {
			return false
		}
	}
	// ensure it contains at least one 'd'
	if !isDieRoll(rollstring) {
		return false
	}

	return true
}

func isOperator(opchar string) bool {
	if !strings.Contains(OPERATORS, string(opchar)) {
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
	// stop at end of of s, record final token, return tokenized slice
	var tokens []string
	var nextchar rune
	last_index := len(rollstring) - 1
	// fmt.Printf("\nrollstring length: %d; last_index: %d", len(rollstring), last_index)
	for i := 0; i <= last_index; i++ {
		if i == last_index {
			// fmt.Printf("Seek got last index. Recording token: ")
			tokens = append(tokens, string(rollstring[i]))
			// fmt.Printf("%s", tokens)
		}
		for j := i + 1; j <= last_index; j++ {
			// fmt.Printf("\ni: %d (%s), j:%d (%s)", i, string(rollstring[i]), j, string(rollstring[j]))
			nextchar = rune(rollstring[j])
			if isOperator(string(nextchar)) {
				// fmt.Printf("got operator %s; recording tokens: ", string(nextchar))
				tokens = append(tokens, string(rollstring[i:j]))
				tokens = append(tokens, string(nextchar))
				// fmt.Printf("%s", tokens)
				// fmt.Printf("\nSetting i to %d", j)
				i = j
				break
			}
		}
	}
	return tokens
}

func ParseRoll(rollstring string) (uint64, uint64, Exception) {
	count_str, sides_str, _ := strings.Cut(rollstring, "d")
	var count uint64
	if count_str != "" {
		count, err := strconv.ParseUint(count_str, 10, 64)
		if err != nil {
			fmt.Printf("failed to convert token %s to uint64", count_str)
			return count, 0, err
		}
	} else {
		count = 1
	}
	sides, err := strconv.ParseUint(sides_str, 10, 64)
	if err != nil {
		fmt.Printf("failed to convert token %s to uint64", count_str)
		return 0, 0, err
	}
	return count, sides, nil
}

// Parse Roll String in format XdN +/- Y and return a list of tokens
func ParseRollString(rollstring string) []string {
	//TODO - implement me
	// counts and sides cannot be zero (modifiers can be even if it's silly)
	// Handle different ordered tokens e.g. mdn+x, x+mdn, mdn+kdn
	// rollstring could result in negative sums; If it does return zero.
	// if it doesn't contain at least one 'd' it isn't valid.

	// Variations:
	// dN - both upper and lower case
	// MdN
	// dN+/-X
	// MdN+/-X
	// MdN+KdN
	// X+/-MdN

	// Algo:
	// Given a slice of token strings
	// For each token:
	// 		If token is not a separator, check for 'd' in string.
	// 		If d is found:
	// 			if d is the first character RollN and return result
	//			if d is NOT the first character RollBatch([0][1]) and return result
	//		If d is NOT found:
	// 			If token is a static int value, record and continue...
	//		If token is a separator, determine which math operation, move to next token, and sum with previous result
	// Finally: return total
	// var separators = [...]rune{'+', '-'}
	// var d rune = 'd'
	// var tokens = [...]string{}
	if !isValidRoll(rollstring) {
		fmt.Printf("Input roll string %s is invalid", rollstring)
		// TODO : raise and handle a sensible exception.
		return []string{}
	}
	rollstring = sanitize(rollstring)
	tokens := tokenize(rollstring)
	return tokens
}

// ===== misc private test functions
func forcedResultRoll(sides uint64, forced uint64) uint64 {
	result := Roll(sides)
	result = forced // override to the value we want.
	return result
}
