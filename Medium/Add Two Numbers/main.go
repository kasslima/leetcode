package main

import "fmt"

func main() {
	//Example
	fmt.Println(addTwoNumbers(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}, &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}))
	fmt.Println(addTwoNumbers(&ListNode{Val: 1, Next: &ListNode{Val: 2}}, &ListNode{Val: 3, Next: &ListNode{Val: 4}}))
}

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    return sumNumbers(l1, l2, 0)
}

func sumNumbers(l1 *ListNode, l2 *ListNode, carry int) *ListNode {

    if l1 == nil && l2 == nil && carry == 0 {
        return nil
    }

    v1 := 0
    v2 := 0

    if l1 != nil {
        v1 = l1.Val
    }

    if l2 != nil {
        v2 = l2.Val
    }

    sum := v1 + v2 + carry

    node := &ListNode{
        Val: sum % 10,
    }

    nextCarry := sum / 10

    var next1 *ListNode
    var next2 *ListNode

    if l1 != nil {
        next1 = l1.Next
    }

    if l2 != nil {
        next2 = l2.Next
    }

    node.Next = sumNumbers(next1, next2, nextCarry)

    return node
}