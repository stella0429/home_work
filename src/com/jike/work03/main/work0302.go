package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	var frontRes *ListNode
	listLength := len(lists)
	if listLength == 0 {
		return frontRes
	}
	if listLength == 1 {
		return lists[0]
	}
	frontRes = lists[0]
	var backRes *ListNode = lists[listLength-1]
	for i := 1; i < listLength; i++ {
		if i >= (listLength+1)/2 {
			break
		}
		frontRes = mergeTwoLists(frontRes, lists[i])
		if listLength%2 == 1 && listLength/2 == i {
			break
		}
		backRes = mergeTwoLists(lists[listLength-1-i], backRes)
	}
	return mergeTwoLists(frontRes, backRes)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func main() {

}
