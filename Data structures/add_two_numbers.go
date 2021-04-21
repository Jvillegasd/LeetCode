// LeetCode #2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    root := &ListNode{}
    for ptr, sum := root, 0; l1 != nil || l2 != nil || sum > 0; sum /= 10 {
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        ptr.Next = &ListNode{
            Val : sum%10,
        }
        ptr = ptr.Next
    }
    return root.Next
}