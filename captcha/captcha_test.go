package captcha_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tatoonz/go-kata/captcha"
)

func TestGenerate_Format(t *testing.T) {
	t.Run("should return left operand as-is and right operand as word when format is 0", func(t *testing.T) {
		assert.Equal(t, captcha.Generate("0", "0", 0, "0"), "0 + zero")
		assert.Equal(t, captcha.Generate("0", "1", 0, "1"), "1 + one")
		assert.Equal(t, captcha.Generate("0", "2", 0, "2"), "2 + two")
		assert.Equal(t, captcha.Generate("0", "3", 0, "3"), "3 + three")
		assert.Equal(t, captcha.Generate("0", "4", 0, "4"), "4 + four")
		assert.Equal(t, captcha.Generate("0", "5", 0, "5"), "5 + five")
		assert.Equal(t, captcha.Generate("0", "6", 0, "6"), "6 + six")
		assert.Equal(t, captcha.Generate("0", "7", 0, "7"), "7 + seven")
		assert.Equal(t, captcha.Generate("0", "8", 0, "8"), "8 + eight")
		assert.Equal(t, captcha.Generate("0", "9", 0, "9"), "9 + nine")
	})

	t.Run("should return left operand as word and right operand as-is when format is 1", func(t *testing.T) {
		assert.Equal(t, captcha.Generate("1", "0", 0, "0"), "zero + 0")
		assert.Equal(t, captcha.Generate("1", "1", 0, "1"), "one + 1")
		assert.Equal(t, captcha.Generate("1", "2", 0, "2"), "two + 2")
		assert.Equal(t, captcha.Generate("1", "3", 0, "3"), "three + 3")
		assert.Equal(t, captcha.Generate("1", "4", 0, "4"), "four + 4")
		assert.Equal(t, captcha.Generate("1", "5", 0, "5"), "five + 5")
		assert.Equal(t, captcha.Generate("1", "6", 0, "6"), "six + 6")
		assert.Equal(t, captcha.Generate("1", "7", 0, "7"), "seven + 7")
		assert.Equal(t, captcha.Generate("1", "8", 0, "8"), "eight + 8")
		assert.Equal(t, captcha.Generate("1", "9", 0, "9"), "nine + 9")
	})
}

func TestGenerate_ConvertOperator(t *testing.T) {
	t.Run("should be + when operator is 0", func(t *testing.T) {
		assert.Equal(t, "1 + one", captcha.Generate("0", "1", 0, "1"))
	})

	t.Run("should be - when operator is 1", func(t *testing.T) {
		assert.Equal(t, "1 - one", captcha.Generate("0", "1", 1, "1"))
	})

	t.Run("should be * when operator is 2", func(t *testing.T) {
		assert.Equal(t, "1 * one", captcha.Generate("0", "1", 2, "1"))
	})
}
