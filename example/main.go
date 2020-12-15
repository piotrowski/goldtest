package example

type Node struct {
	Value    int
	Children []*Node
}

func Sum(x, y int) int {
	return x + y
}

func CreateTree(size int) *Node {
	root := Node{Value: 0, Children: []*Node{{Value: 1}, {Value: 2}}}

	currentNodes := []*Node{root.Children[0], root.Children[1]}
	for i := 3; i <= size; i++ {
		newNodes := []*Node{}
		for _, n := range currentNodes {
			n.Children = []*Node{{Value: i}, {Value: i + 1}}
			newNodes = append(newNodes, n.Children[0])
			newNodes = append(newNodes, n.Children[1])
			i += 2
		}
		currentNodes = newNodes
	}

	return &root
}
