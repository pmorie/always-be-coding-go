package trees

// DepthFirstTraversal returns an array of strings containing the values of the
// given tree in depth-first order.
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
// ...returns: A, B, D, G, C, E, F, H, I
func DepthFirstTraversal(root *Node) []string {
	values := []string{}
	values = append(values, root.Value)
	if root.LNode != nil {
		values = append(values, DepthFirstTraversal(root.LNode)...)
	}
	if root.RNode != nil {
		values = append(values, DepthFirstTraversal(root.RNode)...)
	}

	return values
}
