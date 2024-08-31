//lint:file-ignore U1000 Ignore all unused code
package validparentheses

/*
COMMENT
*/
func isValidStep1(s string) bool {
	var stack []rune
	brackets := map[rune]rune{')': '(', '}': '{', ']': '['}
	for _, bracket := range s {
		if b, ok := brackets[bracket]; !ok {
			stack = append(stack, bracket)
		} else {
			if len(stack) == 0 {
				return false
			}
			if b != stack[len(stack)-1] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}
