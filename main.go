package main

import (
	"arithmetic/link"
)

func main() {

	//head :=link.NewListNode()
	//head.InsertNode(1)
	//head.InsertNode(2)
	//head.InsertNode(3)
	//head.InsertNode(5)
	//head.InsertNode(4)
	//head.InsertNode(5)
	//head.InsertNode(6)
	//head.ListNode()
	//fmt.Println("The Length of List:", head.Length())
	////link.ReverseList(head, head.GetFirstNode())
	////head.GetFirstNode()
	////
	////head.ListNode()
	//head.DelReverseNode(7)
	//head.ListNode()
	head := link.NewChainNode()
	head.CreateChain()
	head.ListChain()
	head.Checkloop()

}
