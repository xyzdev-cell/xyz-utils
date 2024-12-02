package xyz_hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CsharpHash(t *testing.T) {
	assert.Equal(t, int64(372029327), CsharpStringHashV1("3"))
	assert.Equal(t, int64(372029326), CsharpStringHashV1("0"))
	assert.Equal(t, int64(518336165), CsharpStringHashV1("MengHuan02"))
	assert.Equal(t, int64(251214853), CsharpStringHashV1("Ball_sports"))
}
