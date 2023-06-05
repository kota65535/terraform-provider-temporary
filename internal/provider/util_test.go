package provider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsFilePath(t *testing.T) {
	contains, _ := ContainsFilePath("aaa/bbb", "aaa/bbb/ccc")
	assert.True(t, contains)

	contains, _ = ContainsFilePath("aaa/bbb", "aaa/bbb/..")
	assert.False(t, contains)
}
