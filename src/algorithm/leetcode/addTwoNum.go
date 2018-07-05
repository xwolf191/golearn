package leetcode

type ListNode struct {
	    Val int
	    Next *ListNode
	}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    //var val = getNumber(l1)
    getNumber(l1)
    return nil
}

func getNumber(node *ListNode) int{

	var val = node.Val
	var next = node.Next
	if next != nil {
		getNumber(next)
	}
	println(val)
	return val
}