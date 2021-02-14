package regxgen

import (
	"bytes"
	"fmt"
	"math/rand"
)

type Generator struct {
	chars []rune
	lengthMax int
	lengthMin int
}

func (generator *Generator) appendChar(char rune) {
	generator.chars = append(generator.chars, char)
}


func (generator *Generator) generate(config *Config) (string, error) {
	//if generator.chars == nil {
	//	return "", fmt.Errorf("chars array cannot be empty")
	//}
	if generator.lengthMin > generator.lengthMax {
		return "", fmt.Errorf("lengthMin should not be greater that lengthMax")
	}
	if config == nil {
		return "", fmt.Errorf("config is nil")
	}
	if generator.lengthMax == -1 {
		generator.lengthMax = config.repetetionMax
	}
	strLen := rand.Intn(generator.lengthMax - generator.lengthMin) + generator.lengthMin
	nChars := len(generator.chars)
	var result bytes.Buffer
	for ; strLen > 0; strLen-- {
		result.WriteRune(generator.chars[rand.Intn(nChars)])
	}
	return result.String(), nil
}