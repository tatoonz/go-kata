package captcha

import "fmt"

func Generate(format, leftOperand, operator, rightOperand string) string {
	operatorSign := "+"
	if operator == "1" {
		operatorSign = "-"
	} else if operator == "2" {
		operatorSign = "*"
	}

	if format == "1" {
		return fmt.Sprintf("one %s 1", operatorSign)
	}

	return fmt.Sprintf("1 %s one", operatorSign)
}
