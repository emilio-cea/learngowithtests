package iteration

func Repeat(letter string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += letter
	}
	return repeated
}
