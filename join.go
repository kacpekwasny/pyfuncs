package pyfuncs

// Join as arg pass any iteratable var
func Join(s string, iter []string) string {
	retStr := ""
	for _, it := range iter {
		retStr = retStr + it + s
	}
	return retStr[0 : len(retStr)-len(s)]
}
