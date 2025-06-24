package main

import (
    "container/list"
    "fmt"
    "sort"
)

// 1. Two Sum - Optimized
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

// 2. Reverse a Linked List
 type ListNode struct {
    Val  int
    Next *ListNode
}

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

// 3. Valid Parentheses
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

// 4. Merge Intervals
 type Interval struct {
    Start int
    End   int
}

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
            last.End = max(last.End, interval.End)
        } else {
            merged = append(merged, interval)
        }
    }
    return merged
}

// 5. Binary Tree Level Order Traversal
 type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

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
