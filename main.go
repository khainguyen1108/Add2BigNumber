package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// tao file log
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("[main.file.err]", err)
		return
	}

	defer file.Close()

	log.SetOutput(file)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Number1: ")
	a, _ := reader.ReadString('\n')

	fmt.Print("Number2: ")
	b, _ := reader.ReadString('\n')

	a = strings.TrimSpace(a)
	b = strings.TrimSpace(b)

	result := sum(a, b)

	fmt.Println("Result:", result)

	log.Println("Input 1:", a, "Input 2:", b, "Result:", result)
}

type Node struct {
	val  int
	next *Node
}

func buildList(num string) *Node {
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

func add(l1 *Node, l2 *Node) *Node {

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

func listToString(head *Node) string {
	res := []byte{}

	for head != nil {
		res = append([]byte{byte(head.val + '0')}, res...)
		head = head.next
	}
	return string(res)
}

func sum(a string, b string) string {

	l1 := buildList(a)
	l2 := buildList(b)

	result := add(l1, l2)

	return listToString(result)
}
