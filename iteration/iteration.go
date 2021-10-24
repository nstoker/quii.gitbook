package iteration

import "log"

func Repeat(character string, iterations int) string {
	if iterations < 0 {
		log.Fatalf("invalid iteration count passed (%d)", iterations)
	}
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}

	return repeated
}
