package bst

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	if tree == nil {
		return -1
	}

	return closestVal(tree, target, tree.Value)
}

func closestVal(root *BST, target, closest int) int {
	if root == nil {
		return closest
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	if abs(target-root.Value) < abs(target-closest) {
		closest = root.Value
	}

	if target == root.Value {
		return root.Value
	} else if target < root.Value {
		closest = closestVal(root.Left, target, closest)
	} else if target > root.Value {
		closest = closestVal(root.Right, target, closest)
	}

	return closest

}
