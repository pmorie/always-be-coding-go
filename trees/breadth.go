package trees

// BreadthFirstTraversal returns an array of strings containing the values of the
// given tree in breadth-first order.
//
// Given:
//
//       A
//      / \
//     B   C
//    /   / \
//   D   E   F
//  /       / \
// G       H   I
//
// ...returns: A, B, C, D, E, F, G, H, I
func BreadthFirstTraversal(root *Node) []string {
	values := []string{}
	leftToExplore := []*Node{root}

	for len(leftToExplore) > 0 {
		// explore the first node left and add its children to the list of
		node := leftToExplore[0]
		values = append(values, node.Value)
		if node.LNode != nil {
			leftToExplore = append(leftToExplore, node.LNode)
		}
		if node.RNode != nil {
			leftToExplore = append(leftToExplore, node.RNode)
		}
		leftToExplore = leftToExplore[1:]
	}

	return values
}
