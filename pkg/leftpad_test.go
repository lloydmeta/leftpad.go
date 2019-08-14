package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeftpad(t *testing.T) {
	expected := "    Hello world"
	observed := Leftpad("Hello world", 15, ' ')
	assert.Equal(t, expected, observed, "observed should be expected")
}
