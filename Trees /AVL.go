package main

import (
    "fmt"
)

type AVLTree struct {
	root *AvlTreeNode
}

func (e *AVLTree) add(x int) {
	e.root = e.root.add(x)
}

type AvlTreeNode struct {
	height int
	left   *AvlTreeNode
	right  *AvlTreeNode
	value  int
}

func (a *AvlTreeNode) add(x int) *AvlTreeNode {
	if a == nil {
		return &AvlTreeNode{0, nil, nil, x}
	}
	if x > a.value {
		a.right = a.right.add(x)
	} else {
		a.left = a.left.add(x)
	}
	a.recalculateheight()
	return a.rebalance()
}
func (a *AvlTreeNode) rebalance() *AvlTreeNode {
	if a == nil {
		return nil
	}
	a.recalculateheight()
	bfactor := a.left.getheight() - a.right.getheight()
	if bfactor == -2 {
		if a.right.left.getheight() > a.right.right.getheight() {
			a.right = a.right.rotateRight()
		}
		return a.rotateLeft()
	} else if bfactor == 2 {
		if a.left.right.getheight() > a.left.left.getheight() {
			a.left = a.left.rotateLeft()
		}
		return a.rotateRight()
	}
	return a
}
func (a *AvlTreeNode) recalculateheight() {
	a.height = 1 + max(a.right.getheight(), a.left.getheight())
}
func (a *AvlTreeNode) getheight() int {
	if a == nil {
		return 0
	} else if a.height == 0 {
		return 1
	}
	return a.height
}
func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
func (a *AvlTreeNode) rotateLeft() *AvlTreeNode {
	newRoot := a.right
	a.right = newRoot.left
	newRoot.left = a
	a.recalculateheight()
	newRoot.recalculateheight()
	return newRoot
}

func (a *AvlTreeNode) rotateRight() *AvlTreeNode {
	newRoot := a.left
	a.left = newRoot.right
	newRoot.right = a
	a.recalculateheight()
	newRoot.recalculateheight()
	return newRoot
}

func main() {
	t := &AVLTree{nil}
	t.add(1)
	t.add(2)
	t.add(3)
	t.add(4)
	t.add(5)
	t.add(6)
	t.add(7)
	t.add(8)
	t.add(9)

	fmt.Println(t.root.value, t.root.right.value, t.root.left.value, t.root.right.right.right.value)
}
