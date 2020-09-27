package pyfuncs

// Split returns slice of strings split by the char
func Split(str string, match string) []string {
	retStr := []string{}
	indexes := []int{}
	ln := len(match)
	for i := range Range(float64(len(str)) - float64(ln)) {
		j := int(i)
		if str[j:j+ln] == match {
			indexes = append(indexes, j)
		}
	}

	// if no mathches return str
	if len(indexes) == 0 {
		return []string{str}
	}

	lastIndex := 0
	for _, i := range indexes {
		retStr = append(retStr, str[lastIndex:i])
		lastIndex = i + ln
	}
	return append(retStr, str[lastIndex:])
}
