package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Value() string {
	return n.value
}

func (n *Node) Length() int {
	length := 0
	linkedList := n
	for linkedList != nil {
		length += 1
		linkedList = linkedList.Next()
	}

	return length
}

func NewLinkedList(s string) *Node {
	if len(s) == 0 {
		return nil
	}

	headNode := &Node{string(s[0]), nil}
	currentNode := headNode

	for _, v := range s[1:] {
		newNode := &Node{string(v), nil}
		currentNode.next = newNode
		currentNode = newNode
	}

	return headNode
}

func PrintLinkedList(linkedList *Node) {
	for linkedList != nil {
		fmt.Println(linkedList)
		linkedList = linkedList.Next()
	}
}

func ReverseList(linkedList *Node, endPoint int) *Node {

	if endPoint == 1 {
		return linkedList
	}

	first, next, lastNode := linkedList, linkedList, linkedList
	lastNode = nil

	for i := 0; i < endPoint; i++ {
		next = first.Next()
		first.next = lastNode
		lastNode = first
		first = next

	}

	return lastNode

}

// func StitchListsBackTogether(listA, listB, listC *Node, endOfListA int) *Node {
// 	listAFixed := ReverseList(listA, endOfListA)
// 	listStart := listAFixed
// 	for listAFixed.Next() != nil {
// 		listAFixed = listAFixed.Next()
// 	}
// 	listAFixed.next = listB

// 	return listStart

// }

func IsPalindromeFast(linkedList *Node) bool {
	// Uses O(1) space and O(n) time
	// However mutates and destroys list, so need to stitch list backtogether.  Initial implementation StitchListsBackTogether
	length := linkedList.Length()

	endOfListA := length / 2
	endOfListB := length / 2

	if length%2 == 0 {
		endOfListB += 1
	} else {
		endOfListB += 2
	}

	listA, listB := linkedList, linkedList

	for i := 0; i < endOfListB-1; i++ {
		listB = listB.Next()
	}

	listA = ReverseList(listA, endOfListA)

	for listA != nil && listB != nil {
		if listA.Value() != listB.Value() {
			return false
		}
		listA = listA.Next()
		listB = listB.Next()
	}

	return true
}

func IsPalindromeSlow(linkedList *Node) bool {
	//Uses O(1) space and O(n^2) time
	startPalidrome, endPalidrome := linkedList, linkedList

	for endPalidrome.Next() != nil {
		endPalidrome = endPalidrome.Next()
	}

	for startPalidrome != endPalidrome {

		if startPalidrome.Value() != endPalidrome.Value() {
			return false
		}

		if startPalidrome.Next() == endPalidrome {
			return true
		}

		startPalidrome = startPalidrome.Next()

		middlePalidrome := startPalidrome

		for middlePalidrome.Next() != endPalidrome {
			middlePalidrome = middlePalidrome.Next()
		}
		endPalidrome = middlePalidrome

	}

	return true
}

func main() {

	fmt.Println(IsPalindromeSlow(NewLinkedList("ttoott")))
	fmt.Println(IsPalindromeFast(NewLinkedList("ttoott")))
	fmt.Println("")
	fmt.Println(IsPalindromeSlow(NewLinkedList("ttott")))
	fmt.Println(IsPalindromeFast(NewLinkedList("ttott")))
	fmt.Println("")
	fmt.Println(IsPalindromeSlow(NewLinkedList("hello")))
	fmt.Println(IsPalindromeFast(NewLinkedList("hello")))
	fmt.Println("")
	fmt.Println(IsPalindromeSlow(NewLinkedList("ttto")))
	fmt.Println(IsPalindromeFast(NewLinkedList("ttto")))
	fmt.Println("")
	fmt.Println(IsPalindromeSlow(NewLinkedList("tt")))
	fmt.Println(IsPalindromeFast(NewLinkedList("tt")))
	fmt.Println("")
	fmt.Println(IsPalindromeSlow(NewLinkedList("t")))
	fmt.Println(IsPalindromeFast(NewLinkedList("t")))
	fmt.Println("")
	fmt.Println(IsPalindromeSlow(NewLinkedList("tto3tt")))
	fmt.Println(IsPalindromeFast(NewLinkedList("tto3tt")))
	fmt.Println("")

}
