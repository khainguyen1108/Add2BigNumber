package add2num

type MyBigNumber struct{}

type Node struct {
	val  int
	next *Node
}

func (bn MyBigNumber) buildList(num string) *Node {
	var head *Node
	var tail *Node

	for i := len(num) - 1; i >= 0; i-- {
		digit := int(num[i] - '0')
		node := &Node{val: digit}

		if head == nil {
			head = node
			tail = node
		} else {
			tail.next = node
			tail = node
		}
	}

	return head
}

func (bn MyBigNumber) add(l1 *Node, l2 *Node) *Node {

	dummy := &Node{}
	cur := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {

		sum := carry

		if l1 != nil {
			sum += l1.val
			l1 = l1.next
		}

		if l2 != nil {
			sum += l2.val
			l2 = l2.next
		}

		carry = sum / 10

		cur.next = &Node{val: sum % 10}
		cur = cur.next
	}

	return dummy.next
}

func (bn MyBigNumber) listToString(head *Node) string {
	res := []byte{}

	for head != nil {
		res = append([]byte{byte(head.val + '0')}, res...)
		head = head.next
	}
	return string(res)
}

func (bn MyBigNumber) Sum(a string, b string) string {

	l1 := bn.buildList(a)
	l2 := bn.buildList(b)

	result := bn.add(l1, l2)

	return bn.listToString(result)
}
