package factory_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactory(t *testing.T) {

	t.Run("Testing Assert", func(t *testing.T) {
		// assert equality
		assert.Equal(t, 123, 123, "they should be equal")
	})

}
