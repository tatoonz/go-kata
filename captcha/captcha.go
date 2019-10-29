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

func Generate(format int, leftOperand int, operator int, rightOperand string) string {
	if format == 1 {
		return fmt.Sprintf("%s %s %s", operandWords[leftOperand], operatorSigns[operator], rightOperand)
	}

	return fmt.Sprintf("%s %s %s", strconv.Itoa(leftOperand), operatorSigns[operator], operand(rightOperand).Word())
}
