package iteration

func Repeat(repeatedChar string) string {
	var response string
	for i := 0; i < 5; i++ {
		response += repeatedChar
	}
	return response

}
