package binary_tree_sequence

import "fmt"

// To execute Go code, please declare a func main() in a package "main"
/*
 1
/ \
2  3
   /\
  4  2
    / \
   1   1
*/

func run() {

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
	fmt.Println("longest chain: ", result)
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func search(node *Node, chain []int) []int {

	if node == nil {
		return chain
	}

	chain = append(chain, node.value)

	if node.left == nil && node.right == nil {
		return chain
	}

	if node.left != nil && node.left.value < node.value && node.right != nil && node.right.value < node.value {
		return chain
	}

	var leftChain []int
	leftChain = copySlice(leftChain, chain)
	if node.left != nil && node.left.value > node.value {
		leftChain = search(node.left, chain)
	}

	var rightChain []int
	rightChain = copySlice(rightChain, chain)
	if node.right != nil && node.right.value > node.value {
		rightChain = search(node.right, chain)
	}

	if len(leftChain) > len(rightChain) {
		return leftChain
	}

	return rightChain
}

func copySlice(dst, src []int) []int {
	for _, v := range src {
		dst = append(dst, v)
	}
	return dst
}
