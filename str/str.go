package str

func Chr(c rune) string {
	return string(c)
}

func Ord(c string) []rune {
	return []rune(c)
}

func Substr(str string, begin, length int) string {
	lth := len(str)	
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	
	return string(str[begin:end])
}