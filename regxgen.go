package regxgen

import "bytes"

func Generate(pattern string, count int) ([]string, error){
	generators, err := parsePattern(pattern)
	if err != nil {
		return nil, err
	}
	results := make([]string, count)
	for _,_ = range results {
		results = append(results, generateRandomString(generators))
	}
	return results, nil
}

func generateRandomString(generators []Generator) string {
	var result bytes.Buffer
	for _,gen := range generators {
		result.WriteString(gen.generate())
	}
	return result.String()
}