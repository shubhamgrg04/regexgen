package regxgen

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	config := Config{
		RepetetionMax: 10,
		Seed:          5,
	}

	strings,err := Generate("[^aeiouAEIOU0-9]{5}", 1, &config)
	if err != nil {
		t.Errorf("generate failed, got error: %s", err.Error())
	}
	if len(strings) < 1 || strings[0] != "mkYkj" {
		t.Errorf("generate failed")
	}

	strings,err = Generate("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{8}", 1, &config)
	if err != nil {
		t.Errorf("generate failed, got error: %s", err.Error())
	}
	if len(strings) < 1 || strings[0] != "ac903ae5-ac90-ac90-ac903ae5" {
		t.Errorf("generate failed")
	}

	strings,err = Generate("[-+]?[0-9]{1,16}[.][0-9]{1,6}", 1, &config)
	if err != nil {
		t.Errorf("generate failed, got error: %s", err.Error())
	}
	if len(strings) < 1 || strings[0] != "-6907669.69" {
		t.Errorf("generate failed")
	}

}