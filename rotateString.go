package rotateString

/*
TODO: Essa não tem um bom desempenho ela realoca muita memória, reescrever usando um buffer.
*/
func Rotate(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
