package bst

import "testing"

func TestBST_FindClosestValue(t *testing.T) {
	type fields struct {
		Value int
		Left  *BST
		Right *BST
	}
	type args struct {
		target int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BST{
				Value: tt.fields.Value,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := tree.FindClosestValue(tt.args.target); got != tt.want {
				t.Errorf("BST.FindClosestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
