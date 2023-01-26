package iteration

var repeatCount = 5

func Repeat(character string, repetitions int) string {
	var repeated string

	if repetitions == 0 {
		repetitions = repeatCount
	}

	for i := 0; i < repetitions; i++ {
		repeated += character
	}
	return repeated
}
