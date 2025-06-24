package main

import (
	"container/list"
	"fmt"
	"sort"
	"testQuestions/golang/islandFinder"
)

// twoSum finds two numbers in an array that add up to the target
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		if index, found := numMap[target-num]; found {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return nil
}

// ListNode represents a node in a singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseLinkedList reverses a singly linked list
func reverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

// isValidParentheses checks if a string has valid matching parentheses
func isValidParentheses(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{')': '(', '}': '{', ']': '['}
	for _, char := range s {
		if open, exists := mapping[char]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

// Interval represents a range with a start and end value
type Interval struct {
	Start int
	End   int
}

// mergeIntervals merges overlapping intervals in a list
func mergeIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	merged := []Interval{intervals[0]}
	for _, interval := range intervals[1:] {
		last := &merged[len(merged)-1]
		if interval.Start <= last.End {
			if interval.End > last.End {
				last.End = interval.End
			}
		} else {
			merged = append(merged, interval)
		}
	}
	return merged
}

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrderTraversal performs a level order traversal on a binary tree
func levelOrderTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		level := []int{}
		size := queue.Len()

		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			level = append(level, node.Val)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		result = append(result, level)
	}
	return result
}

// main function to demonstrate the implemented functions
func main() {
	fmt.Println("Running Go implementations...")

	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println("Two Sum:", twoSum(nums, target))

	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	fmt.Println("Reversed Linked List:", reverseLinkedList(head))

	fmt.Println("Valid Parentheses:", isValidParentheses("{[()]}"))

	intervals := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println("Merged Intervals:", mergeIntervals(intervals))

	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println("Binary Tree Level Order Traversal:", levelOrderTraversal(root))

	grid := [][]int{
		{0, 1, 1, 0, 0},
		{1, 0, 1, 0, 1},
		{0, 0, 0, 1, 1},
		{1, 0, 0, 0, 0},
	}

	fmt.Printf("Island Count: %d\n", islandFinder.CountIslands(grid))
}
