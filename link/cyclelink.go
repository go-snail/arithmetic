package link

import (
	"fmt"
)

type chainNode struct {
	num  int
	next *chainNode
}

var gc = make(map[*chainNode]int, 0)

func NewChainNode() *chainNode {
	return &chainNode{0, nil}
}
func (l *chainNode) ListChain() {
	for l.next != nil && gc[l.next] == 0 {
		gc[l.next] = 1
		fmt.Printf(" %d", l.next.num)
		l = l.next
	}
}

func (pHead *chainNode) CreateChain() {
	var temp *chainNode
	var a *chainNode
	for i := 1; i <= 7; i++ {
		temp = NewChainNode()
		temp.num = i
		pHead.next = temp
		pHead = pHead.next
		if i == 4 {
			a = temp
		}
	}
	temp.next = a
}

func (pHead *chainNode)Checkloop()  {
	var fast =  pHead.next
	var slow = pHead.next
	slow = slow.next
	fast = fast.next.next

	for slow != fast && fast != nil && slow != nil {
		fast = fast.next.next
		slow = slow.next
	}
	fmt.Println("交点b:", fast.num)
	slow = slow.next
	len1 := 1
	//fast不动，slow继续走再次相等则为环的长度
	for slow != fast {
		slow = slow.next
		len1++
	}
	fmt.Println("This length of cycle is:", len1)

	//求长度
	len2 := 0
	slow = pHead.next
	for slow != fast {
		fast = fast.next
		slow = slow.next
		len2 ++
	}
	fmt.Println("This length of chain is:", len1 + len2)
}