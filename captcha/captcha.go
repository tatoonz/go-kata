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

func Generate(format, leftOperand, operator, rightOperand string) string {
	if format == "1" {
		i, _ := strconv.Atoi(leftOperand)
		return fmt.Sprintf("%s %s %s", operandWords[i], operatorSignMap[operator], rightOperand)
	}

	i, _ := strconv.Atoi(rightOperand)
	return fmt.Sprintf("%s %s %s", leftOperand, operatorSignMap[operator], operandWords[i])
}
