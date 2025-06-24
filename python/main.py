from collections import deque
import islandfinder

### 1. Two Sum (Array & HashMap) - Python
def two_sum(nums, target):
    num_map = {}
    for i, num in enumerate(nums):
        complement = target - num
        if complement in num_map:
            return [num_map[complement], i]
        num_map[num] = i
    return []

### 2. Reverse a Linked List - Python
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def reverse_linked_list(head):
    prev = None
    current = head
    while current:
        next_node = current.next
        current.next = prev
        prev = current
        current = next_node
    return prev

### 3. Valid Parentheses - Python
def is_valid_parentheses(s):
    stack = []
    mapping = {')': '(', '}': '{', ']': '['}
    for char in s:
        if char in mapping:
            top_element = stack.pop() if stack else '#'
            if mapping[char] != top_element:
                return False
        else:
            stack.append(char)
    return not stack

### 4. Merge Intervals - Python
def merge_intervals(intervals):
    if not intervals:
        return []
    intervals.sort(key=lambda x: x[0])
    merged = [intervals[0]]
    for interval in intervals[1:]:
        if interval[0] <= merged[-1][1]:
            merged[-1][1] = max(merged[-1][1], interval[1])
        else:
            merged.append(interval)
    return merged

### 5. Binary Tree Level Order Traversal - Python
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def level_order_traversal(root):
    if not root:
        return []
    result = []
    queue = deque([root])
    while queue:
        level = []
        for _ in range(len(queue)):
            node = queue.popleft()
            level.append(node.val)
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        result.append(level)
    return result

### **Main function to test implementations**
def main():
    # Test Two Sum
    nums = [2, 7, 11, 15]
    target = 9
    print(f"Two Sum: {two_sum(nums, target)}")

    # Test Reverse Linked List
    head = ListNode(1, ListNode(2, ListNode(3)))
    reversed_head = reverse_linked_list(head)
    reversed_list = []
    while reversed_head:
        reversed_list.append(reversed_head.val)
        reversed_head = reversed_head.next
    print(f"Reversed Linked List: {reversed_list}")

    # Test Valid Parentheses
    print(f"Valid Parentheses: {is_valid_parentheses('{[()]}')}")

    # Test Merge Intervals
    intervals = [[1, 3], [2, 6], [8, 10], [15, 18]]
    print(f"Merged Intervals: {merge_intervals(intervals)}")

    # Test Binary Tree Level Order Traversal
    root = TreeNode(1, TreeNode(2), TreeNode(3))
    print(f"Binary Tree Level Order Traversal: {level_order_traversal(root)}")
    
    """Main function to test the island finder module."""
    grid = [
        [0, 1, 1, 0, 0],
        [1, 0, 1, 0, 1],
        [0, 0, 0, 1, 1],
        [1, 0, 0, 0, 0]
    ]

    print(f"Island Count: {islandfinder.count_islands(grid)}")

# Ensure the script runs only if executed directly
if __name__ == "__maian__":
    main()
