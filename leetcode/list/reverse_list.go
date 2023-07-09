package list

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func ReverseList(node *Node) *Node {
	var p *Node = nil
	var n *Node = nil
	t := node
	for t != nil {
		// 记录 当前节点的下一个节点
		n = t.Next
		// 当前节点的下一个节点指向前一个节点
		t.Next = p
		// 记录 当前节点为反转列表当前最后一个节点
		p = t
		// 将下一个节点赋值给 t，进入下一个循环
		t = n
	}
	return p
}

func DisplayLinkedList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
	fmt.Println()
}
