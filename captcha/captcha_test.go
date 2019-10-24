package captcha_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tatoonz/go-kata/captcha"
)

func TestGenerate(t *testing.T) {
	t.Run("should return left operand as-is and right operand as text when format is 0", func(t *testing.T) {
		assert.Equal(t, "1 + one", captcha.Generate("0", "1", "0", "1"))
	})

	t.Run("should return left operand as text and right operand as-is when format is 1", func(t *testing.T) {
		assert.Equal(t, "one + 1", captcha.Generate("1", "1", "0", "1"))
	})
}
