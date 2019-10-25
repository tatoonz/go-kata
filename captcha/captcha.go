package captcha

import (
	"fmt"
)

var operatorSignMap = map[string]string{
	"0": "+",
	"1": "-",
	"2": "*",
}

func Generate(format, leftOperand, operator, rightOperand string) string {
	if format == "1" {
		return fmt.Sprintf("one %s 1", operatorSignMap[operator])
	}

	return fmt.Sprintf("1 %s one", operatorSignMap[operator])
}
