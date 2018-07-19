package trees

import (
	"reflect"
	"testing"
)

func TestDepthFirstTraversal(t *testing.T) {
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
			//     B   D
			//    /   / \
			//   C   E   F
			tree: &Node{
				Value: "A",
				LNode: &Node{
					Value: "B",
					LNode: &Node{
						Value: "C",
					},
				},
				RNode: &Node{
					Value: "D",
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
		if actual, expected := DepthFirstTraversal(cases[i].tree), cases[i].expected; !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v: expected %v, got %v", i, expected, actual)
		}
	}
}
