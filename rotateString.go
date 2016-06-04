package rotateString

// Rotate string
func Rotate(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
