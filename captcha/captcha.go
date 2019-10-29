package captcha

import (
	"fmt"
	"strconv"
)

var (
	operatorSigns = []string{"+", "-", "*"}
	operandWords  = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

type operand string

func (op operand) Word() string {
	i, _ := strconv.Atoi(string(op))
	return operandWords[i]
}

func Generate(format, leftOperand string, operator int, rightOperand string) string {
	if format == "1" {
		return fmt.Sprintf("%s %s %s", operand(leftOperand).Word(), operatorSigns[operator], rightOperand)
	}

	return fmt.Sprintf("%s %s %s", leftOperand, operatorSigns[operator], operand(rightOperand).Word())
}
