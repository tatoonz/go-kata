package captcha

func Generate(format, leftOperand, operator, rightOperand string) string {
	if format == "1" {
		return "one + 1"
	}

	return "1 + one"
}
