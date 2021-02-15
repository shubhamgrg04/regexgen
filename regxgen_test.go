package regxgen

import (
	"math/rand"
	"testing"
)

func TestGenerate(t *testing.T) {
	rand.Seed(10)
	gen := Generator{
		chars:     nil,
		lengthMax: 1,
		lengthMin: 5,
	}
	if _,err := gen.generate(&DEFAULT_CONFIG); err == nil {
		t.Errorf("generate failed, expected error, got nil")
	}
}