package trees

import (
	"reflect"
	"testing"
)

func TestBreadthFirstTraversal(t *testing.T) {
	cases := []struct {
		tree     *Node
		expected []string
	}{
		{
			//   A
			//  / \
			// B   C
			tree: &Node{
				Value: "A",
				LNode: &Node{
					Value: "B",
				},
				RNode: &Node{
					Value: "C",
				},
			},
			expected: []string{"A", "B", "C"},
		},
		{
			//       A
			//      / \
			//     B   C
			//    /   / \
			//   D   E   F
			tree: &Node{
				Value: "A",
				LNode: &Node{
					Value: "B",
					LNode: &Node{
						Value: "D",
					},
				},
				RNode: &Node{
					Value: "C",
					LNode: &Node{
						Value: "E",
					},
					RNode: &Node{
						Value: "F",
					},
				},
			},
			expected: []string{"A", "B", "C", "D", "E", "F"},
		},
	}

	for i, _ := range cases {
		tree := cases[i].tree
		if actual, expected := BreadthFirstTraversal(tree), cases[i].expected; !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v: expected %v, got %v", i, expected, actual)
		}
	}
}
