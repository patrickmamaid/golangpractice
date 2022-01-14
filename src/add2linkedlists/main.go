package main


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

// insert
func (l *ListNode) insert(ival int) {
	l.Val = ival
	l.Next = new(ListNode)
}


//// array to decimal
//func arrayToDecimal() int {
//
//}


// decimal to array



func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// break out the ordered list assume max 100
    // 111



	return nil
}






func main() {
	//[2,4,3]
	//[5,6,4]
	//Explanation: 342 + 465 = 807

	// basic operations first
	// insert and print

	var testlist = new(ListNode)
	testlist.insert(1)

}
