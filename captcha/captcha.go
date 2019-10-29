package captcha

import (
	"fmt"
)

var (
	operatorSigns = []string{"+", "-", "*"}
	operandWords  = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

func Generate(format int, leftOperand int, operator int, rightOperand int) string {
	if format == 1 {
		return fmt.Sprintf("%s %s %d", operandWords[leftOperand], operatorSigns[operator], rightOperand)
	}

	return fmt.Sprintf("%d %s %s", leftOperand, operatorSigns[operator], operandWords[rightOperand])
}
