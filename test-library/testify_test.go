package testlibrary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	t.Run("Test nil", func(t *testing.T) {
		var p *struct{} = nil
		assert.Nil(t, p)
	})

	t.Run("Test not nil", func(t *testing.T) {
		object := "Hello world"
		assert.NotNil(t, object)
	})

	t.Run("Test equal", func(t *testing.T) {
		assert.Equal(t, 123, 123, "they should be equal")
	})

	t.Run("Test not equal", func(t *testing.T) {
		assert.NotEqual(t, 123, 321, "they should not be equal")
	})
}
