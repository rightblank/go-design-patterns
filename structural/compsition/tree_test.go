package compsition

import "testing"

func TestTree(t *testing.T) {

    t.Log("LOG: testing tree")
    root := Tree {
        LeafValue:5,
        Left: &Tree{
            LeafValue:4,
            Right: &Tree{3, nil, nil},
        },
    }
    println(root.Left.Right.LeafValue)
}
