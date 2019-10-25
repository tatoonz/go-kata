package captcha

import (
	"fmt"
	"strconv"
)

var (
	operatorSignMap = map[string]string{
		"0": "+",
		"1": "-",
		"2": "*",
	}

	operandWords = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

type operand string

func (op operand) Word() string {
	i, _ := strconv.Atoi(string(op))
	return operandWords[i]
}

func Generate(format, leftOperand, operator, rightOperand string) string {
	if format == "1" {
		return fmt.Sprintf("%s %s %s", operand(leftOperand).Word(), operatorSignMap[operator], rightOperand)
	}

	return fmt.Sprintf("%s %s %s", leftOperand, operatorSignMap[operator], operand(rightOperand).Word())
}
