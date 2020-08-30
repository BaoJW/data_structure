package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibByIteration(t *testing.T) {

	assert.Equal(t, 0, FibByIteration(0))
	assert.Equal(t, 1, FibByIteration(1))
	assert.Equal(t, 1, FibByIteration(2))
	assert.Equal(t, 2, FibByIteration(3))
	assert.Equal(t, 3, FibByIteration(4))
	assert.Equal(t, 5, FibByIteration(5))
	assert.Equal(t, 8, FibByIteration(6))
	assert.Equal(t, 13, FibByIteration(7))
	assert.Equal(t, 21, FibByIteration(8))
	assert.Equal(t, 34, FibByIteration(9))
	assert.Equal(t, 55, FibByIteration(10))
	assert.Equal(t, 89, FibByIteration(11))
	assert.Equal(t, 144, FibByIteration(12))
}

func TestFibByRecursion(t *testing.T) {

	assert.Equal(t, 0, FibByRecursion(0))
	assert.Equal(t, 1, FibByRecursion(1))
	assert.Equal(t, 1, FibByRecursion(2))
	assert.Equal(t, 2, FibByRecursion(3))
	assert.Equal(t, 3, FibByRecursion(4))
	assert.Equal(t, 5, FibByRecursion(5))
	assert.Equal(t, 8, FibByRecursion(6))
	assert.Equal(t, 13, FibByRecursion(7))
	assert.Equal(t, 21, FibByRecursion(8))
	assert.Equal(t, 34, FibByRecursion(9))
	assert.Equal(t, 55, FibByRecursion(10))
	assert.Equal(t, 89, FibByRecursion(11))
	assert.Equal(t, 144, FibByRecursion(12))
}
