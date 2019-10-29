package captcha

import (
	"fmt"
	"strconv"
)

var (
	operatorSigns = []string{"+", "-", "*"}
	operandWords  = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

func Generate(format int, leftOperand int, operator int, rightOperand int) string {
	if format == 1 {
		return fmt.Sprintf("%s %s %s", operandWords[leftOperand], operatorSigns[operator], strconv.Itoa(rightOperand))
	}

	return fmt.Sprintf("%s %s %s", strconv.Itoa(leftOperand), operatorSigns[operator], operandWords[rightOperand])
}
