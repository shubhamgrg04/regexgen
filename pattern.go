package regxgen

import (
	"fmt"
	"strconv"
	"strings"
)

type Pattern struct {
	regex []rune
	currentIndex int
}

func NewPattern(regex string) Pattern{
	return Pattern{
		regex:        []rune(regex),
		currentIndex: 0,
	}
}

// returns rune at the currentIndex
func (pt *Pattern) getCurrentRune() rune{
	return pt.regex[pt.currentIndex]
}

//	helper functions

// creates a []rune with runes from start to end
func makeRange(start, end rune) (result []rune){
	for i:=start; i<=end; i++ {
		result = append(result, i)
	}
	return result
}
// concatenates two []rune into a single []rune
func concatRange(ranges ...[]rune) (result []rune) {
	for _,r := range ranges {
		result = append(result, r...)
	}
	return result
}
// removes negated []rune from []parent
func negateRange(parent, negated []rune) (result []rune) {
	negMap := make(map[rune]bool)
	for _,r := range negated {
		negMap[r] = true
	}
	for _,r := range parent {
		if _,ok := negMap[r]; !ok {
			result = append(result, r)
		}
	}
	return result
}
var allChar = concatRange(makeRange('A','Z'), makeRange('a','z'), makeRange('0','9'))

// 	Returns an array of generators.
// 	Each generator one single Char Class
// 	ex: /[a-b]*[1-4]{1,10}/ will create two generators
// 	one of [a-b]* and the other one for [1-4]{1,10}
func (pt *Pattern)createGenerators() ([]Generator, error){
	var generators []Generator
	lastGenerator := func() *Generator {
		return &generators[len(generators) - 1]
	}
	shouldNegate := false
	pt.currentIndex = 0
	for pt.currentIndex < len(pt.regex) {
		char := pt.getCurrentRune()
		pt.currentIndex++

		switch char {
		case '.':
			generators = append(generators, NewGenerator(allChar))
		case '[':
			g, err := pt.parseBracket()
			if err != nil {
				return nil, err
			}
			generators = append(generators, *g)
		case '^':
			shouldNegate = true
		case '*':
			lastGenerator().setLengthLimits(0, INFINITE)
		case '+':
			lastGenerator().setLengthLimits(1, INFINITE)
		case '?':
			lastGenerator().setLengthLimits(0,1)
		case '{':
			if len(generators) < 1 {
				return generators, fmt.Errorf("invalid braces")
			}
			min, max, err := pt.parseBrace()
			if err != nil {
				return generators, err
			}
			lastGenerator().setLengthLimits(min,max)
		default:
			if shouldNegate {
				generators = append(generators, NewGenerator(negateRange(allChar, []rune{char})))
			} else {
				generators = append(generators, NewGenerator([]rune{char}))
			}
		}
	}
	return generators, nil
}

// parses char classes contained in bracket
func (pt *Pattern) parseBracket() (*Generator, error) {
	var charRange []rune
	escaped := false
	shouldNegate := false
	rangeJustFinished := false
	var startChar rune

	if pt.getCurrentRune() == '^' {
		shouldNegate = true
		pt.currentIndex++
	}
	for {
		if len(pt.regex)-pt.currentIndex < 1 {
			return nil, fmt.Errorf("unclosed []")
		}
		char := pt.getCurrentRune()
		pt.currentIndex++

		if !escaped {

			if char == ']' {
				break
			}

			if char == '\\' {
				escaped = true
				continue
			}

			if char == '-' && len(charRange) > 0 {
				if rangeJustFinished {
					return nil, fmt.Errorf("invalid range")
				}
				startChar = charRange[len(charRange)-1]
				charRange = charRange[:len(charRange)-1]
				continue
			}
		}

		escaped = false
		if startChar != 0 {
			r := makeRange(startChar, char)
			charRange = append(charRange, r...)
			rangeJustFinished = true
			startChar = 0
			continue
		}
		rangeJustFinished = false
		charRange = append(charRange, char)
	}
	var gen Generator
	if shouldNegate {
		gen = NewGenerator(negateRange(allChar, charRange))
	} else {
		gen = NewGenerator(charRange)
	}
	return &gen, nil
}

// parses brace based ranges in pattern
func (pt *Pattern) parseBrace() (int, int, error) {
	hasComma := false
	var buf []rune
	for {
		if len(pt.regex)-pt.currentIndex < 1 {
			return 0,0,fmt.Errorf("bad range %s", string(buf))
		}
		char := pt.getCurrentRune()
		pt.currentIndex++
		if char == '}' {
			break
		}
		if char == ',' {
			hasComma = true
		}
		buf = append(buf, char)
	}
	var min, max int

	if hasComma {
		limits := strings.Split(string(buf), ",")
		if len(limits) !=2 {
			return 0,0,fmt.Errorf("invalid range: {%s}", string(buf))
		}
		var err error
		min,err = strconv.Atoi(limits[0])
		if err != nil {
			return 0,0,err
		}
		if limits[1] == "" {
			max = INFINITE
		} else {
			max,err = strconv.Atoi(limits[1])
			if err != nil {
				return 0,0,err
			}
		}
		if max < min {
			return 0,0,fmt.Errorf("bad range {%s}", string(buf))
		}
	} else {
		min, _ = strconv.Atoi(string(buf))
		max = min
	}
	return min,max,nil
}
