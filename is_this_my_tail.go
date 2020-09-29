package kata

func CorrectTail(body string, tail rune) bool {
	return rune(body[len(body)-1]) == tail
}

