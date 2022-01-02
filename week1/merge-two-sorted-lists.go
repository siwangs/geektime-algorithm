/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{}
	pointer := prehead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pointer.Next = l1
			l1 = l1.Next
		} else {
			pointer.Next = l2
			l2 = l2.Next
		}
		pointer = pointer.Next
	}
	if l1 == nil {
		pointer.Next = l2
	} else {
		pointer.Next = l1
	}
	return prehead.Next
}