package link

import (
	"fmt"
	"github.com/pkg/errors"
)

type ListNode struct {
	data int
	next *ListNode
}

func NewListNode() *ListNode {
	return &ListNode{next: nil}
}

func (l *ListNode) Length() int {
	len := 0
	for l.next != nil {
		len++
		l = l.next
	}
	return len
}

func (l *ListNode) InsertNode(d int) error {
	newNode := new(ListNode)
	newNode.data = d
	newNode.next = l.next
	l.next = newNode
	return nil
}

func (l *ListNode) DelNode(d int) {
	if l == nil {
		errors.New("Empty List!")
		return
	}
	for l.next != nil {
		if l.next.data == d {
			l.next = l.next.next
			//return  此处控制找到相同数据是否全部删除操作
		}
		l = l.next
	}
}

func (l *ListNode)DelReverseNode(d int)  {
	len := l.Length()
	l.DelNodeById(len-d+1)
}

func (l *ListNode) ListNode() {
	for l.next != nil {
		fmt.Printf(" %d", l.next.data)
		l = l.next
	}
}

func (l *ListNode) GetFirstNode() *ListNode {
	return l.next
}
func ReverseList(pHead, node *ListNode) *ListNode {

	if node.next == nil {
		pHead.next = node
		return node
	}
	n := ReverseList(pHead, node.next)
	if n != nil {
		n.next = node
		node.next = nil
	}
	return node
}

func (pHead *ListNode) ReverseListV2() {
	pReversedHead := pHead
	var pNode = pHead.next
	var pPrev *ListNode

	for pNode != nil {
		pNext := pNode.next
		if pNext == nil {
			pReversedHead.next = pNode
		}
		pNode.next = pPrev
		pPrev = pNode
		pNode = pNext
	}

	return
}

func (pHead * ListNode)DelNodeById(id int)  {
	i := 1
	node := pHead.next
	tmp := pHead
	for node != nil {
		if i == id {
			tmp.next = node.next
		}
		i++
		tmp = node
		node = node.next

	}
	if i < id {
		errors.New("The Len of the list is less!")
	}

}