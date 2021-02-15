package regxgen

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Generator struct {
	chars []rune
	lengthMax int
	lengthMin int
}

func NewGenerator(chars []rune) Generator {
	return Generator{
		chars:     chars,
		lengthMax: 1,
		lengthMin: 1,
	}
}

func (generator *Generator) appendChar(char rune) {
	generator.chars = append(generator.chars, char)
}

func (generator *Generator) setLengthLimits(min,max int) {
	generator.lengthMin, generator.lengthMax = min, max
}

func (generator *Generator) generate(config *Config) (string, error) {
	if generator.lengthMax == INFINITE {
		generator.lengthMax = config.RepetetionMax
	}
	if generator.lengthMin > generator.lengthMax {
		return "", fmt.Errorf("lengthMin should not be greater that lengthMax")
	}
	if config == nil {
		return "", fmt.Errorf("config can't be nil while generating strings")
	}

	strLen := generator.lengthMin
	if generator.lengthMax > generator.lengthMin {
		strLen = strLen + rand.Intn(generator.lengthMax - generator.lengthMin)
		if generator.lengthMin == 0 && generator.lengthMax == 1 {
			strLen++
		}
	}

	nChars := len(generator.chars)
	var result bytes.Buffer
	for ; strLen > 0; strLen-- {
		result.WriteRune(generator.chars[rand.Intn(nChars)])
	}
	return result.String(), nil
}