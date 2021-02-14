package regxgen

import "testing"

func TestGenerate(t *testing.T) {
	gen := Generator{
		chars:     nil,
		lengthMax: 1,
		lengthMin: 5,
	}
	if _,err := gen.generate(&DEFAULT_CONFIG); err == nil {
		t.Errorf("generate failed, expected error, got nil")
	}
}