package captcha

import (
	"fmt"
)

var (
	formats       = []string{"%d %s %s", "%s %s %d"}
	operatorSigns = []string{"+", "-", "*"}
	operandWords  = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

type operand int

func (o operand) String() string {
	return operandWords[o]
}

/*
Generate returns captcha in string format.

Operator can be 0, 1, or 2 represents +, - and * respectively.

Format can be:
  - 0 the leftOperand will show the number as-is but the rightOperand will show the number in word
  - 1 the leftOperand will show the number in word  but the rightOperand will show the number as-is
*/
func Generate(format int, leftOperand int, operator int, rightOperand int) string {
	return fmt.Sprintf(formats[format], operand(leftOperand), operatorSigns[operator], operand(rightOperand))
}
