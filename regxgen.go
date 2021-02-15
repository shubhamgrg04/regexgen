package regxgen

import "bytes"

func Generate(regex string, count int, config *Config) ([]string, error){
	pattern := NewPattern(regex)
	generators, err := pattern.createGenerators()
	if err != nil {
		return nil, err
	}
	if config == nil {
		config = &DEFAULT_CONFIG
	}
	var results []string
	for ; count>0; count-- {
		if str, err := generateRandomString(generators, config); err == nil {
			results = append(results, str)
		} else {
			return nil, err
		}
	}
	return results, nil
}

func generateRandomString(generators []Generator, config *Config) (string, error) {
	var result bytes.Buffer
	for _,gen := range generators {
		if str, err := gen.generate(config); err == nil {
			result.WriteString(str)
		} else {
			return "", err
		}
	}
	return result.String(), nil
}