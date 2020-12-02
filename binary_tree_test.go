package binary_tree_sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequence(t *testing.T) {
	root := &Node{
		value: 1,
		left: &Node{
			value: 2,
		},
		right: &Node{
			value: 3,
			left: &Node{
				value: 4,
			},
			right: &Node{
				value: 2,
				left: &Node{
					value: 1,
				},
				right: &Node{
					value: 1,
				},
			},
		},
	}
	result := search(root, []int{})
	assert.NotEmpty(t, result)
	assert.Equal(t, []int{1, 3, 4}, result)
}

func TestSequenceRootEmpty(t *testing.T) {

	var root *Node
	result := search(root, []int{})
	assert.Empty(t, result)
}

func TestSequenceRootEmptyChildren(t *testing.T) {
	root := &Node{
		value: 1,
	}
	result := search(root, []int{})
	assert.NotEmpty(t, result)
	assert.Equal(t, []int{1}, result)
}

func TestSequenceMultipleLongestChain(t *testing.T) {
	root := &Node{
		value: 1,
		left: &Node{
			value: 2,
		},
		right: &Node{
			value: 3,
			left: &Node{
				value: 4,
			},
			right: &Node{
				value: 5,
				left: &Node{
					value: 6,
				},
				right: &Node{
					value: 10,
				},
			},
		},
	}
	result := search(root, []int{})
	assert.NotEmpty(t, result)
	assert.Equal(t, []int{1, 3, 5, 10}, result)
}
