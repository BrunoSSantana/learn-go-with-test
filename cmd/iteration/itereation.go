package iteration

func Repeat(value string, repeatCount int) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += value
	}
	return
}
