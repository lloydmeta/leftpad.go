package pkg

func Leftpad(s string, desiredLength uint32, c rune) string {
	lengthToPad := int(desiredLength) - len(s)
	toPad := s
	for i := 0; i < lengthToPad; i++ {
		toPad = string(c) + toPad
	}
	return toPad
}
