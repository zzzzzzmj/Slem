package leetcode0147

import (
	"github.com/Ackerr/Algorithms/utils"
	"reflect"
	"testing"
)

func TestInsertionSortList(t *testing.T) {
	a := utils.ListNode{Val: 3}
	b := utils.ListNode{Val: 4, Next: &a}
	c := utils.ListNode{Val: 5, Next: &b}
	d := utils.ListNode{Val: 6, Next: &c}
	sortLinkList := insertionSortList(&d)
	result := make([]int, 0)
	for sortLinkList != nil {
		result = append(result, sortLinkList.Val)
		sortLinkList = sortLinkList.Next
	}
	correct := []int{3, 4, 5, 6}
	if !reflect.DeepEqual(result, correct) {
		t.Errorf("result error, correct result should be %d", correct)

	}

}
