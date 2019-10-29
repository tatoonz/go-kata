package captcha

import (
	"fmt"
)

var (
	operatorSigns = []string{"+", "-", "*"}
	operandWords  = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

/*
Generate returns captcha in string format.

Operator can be 0, 1, or 2 represents +, - and * respectively.

Format can be:
  - 0 the leftOperand will show the number as-is but the rightOperand will show the number in word
  - 1 the leftOperand will show the number in word  but the rightOperand will show the number as-is
*/
func Generate(format int, leftOperand int, operator int, rightOperand int) string {
	if format == 1 {
		return fmt.Sprintf("%s %s %d", operandWords[leftOperand], operatorSigns[operator], rightOperand)
	}

	return fmt.Sprintf("%d %s %s", leftOperand, operatorSigns[operator], operandWords[rightOperand])
}
