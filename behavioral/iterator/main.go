package main

import "fmt"

type Iterator interface {
	Next() bool
	Scan() (interface{}, error)
}

type Node struct {
	val int
	left *Node
	right *Node
}

func (n *Node) String() string {
	if n == nil {
		return "nil "
	}
	return n.right.String() + fmt.Sprintf("%d ", n.val) + n.left.String()
}

type BinaryIterator struct {
	Queue []*Node
}

func (b *BinaryIterator) Next() bool {
	if len(b.Queue) == 0 {
		return false
	}
	return true
}

func (b *BinaryIterator) Scan() (interface{}, error) {
	if ok := b.Next(); !ok {
		return &Node{}, nil
	}

	//fmt.Println("1. Length:", len(b.Queue))

	var node = b.Queue[0]
	if node == nil {
		return nil, nil
	}
	if node.left != nil {
		b.Queue = append(b.Queue, node.left)
	}
	if node.right != nil {
		b.Queue = append(b.Queue, node.right)
	}

	temp := b.Queue[1:]
	b.Queue = temp

	//fmt.Println("2. Length:", len(b.Queue))

	return node, nil
}

func main() {
	root := Node{
		val:   5,
		left:  &Node{
			val:   3,
			left:  &Node{
				val:   1,
				left:  nil,
				right: nil,
			},
			right: nil,
		},
		right: &Node{
			val:   10,
			left:  nil,
			right: nil,
		},
	}

	//fmt.Println(root.String())

	var bIterator Iterator = &BinaryIterator{
		Queue: []*Node{
			&root,
		},
	}

	for bIterator.Next() {
		node, _ := bIterator.Scan()
		fmt.Printf("%v \n", node.(*Node).val)
	}
}
