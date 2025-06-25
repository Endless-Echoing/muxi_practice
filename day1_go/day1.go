package main

import "fmt"

func main() {
	n := 30
	primes := Prime(n)
	println("练习1:")
	for _, prime := range primes {
		println(prime)
	}

	println("\n练习2:")
	nums := []int{1, 2, 2, 3, -1, 4, 4, 5}
	unique := Deduplicate(nums)
	for _, num := range unique {
		println(num)
	}

	println("\n练习3:")
	root := BuildTree(nums)
	println(root.val)

	println("\n练习4:")
	nums1 := []int{1, 3, 4, 5, 0, 0}
	m := 4
	nums2 := []int{2, 5, 6}
	n = 3
	nums1 = append(nums1[:m], make([]int, n)...) //nums1长度不够，所以扩展了
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

}

func Prime(n int) []int {
	if n < 2 {
		return []int{}
	} else if n == 2 {
		return []int{2}
	}
	primes := []int{2}
	for i := 3; i <= n; i += 2 {
		isPrime := true
		for _, p := range primes {
			if p*p > i {
				break
			}
			if i%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}

func Deduplicate(nums []int) []int {
	unique := []int{}
	for _, num := range nums {
		exist := true
		for _, value := range unique {
			if num == value {
				exist = false
				break
			}
		}
		if exist {
			unique = append(unique, num)
		}
	}
	return unique
}

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// 传入数组，返回数的根节点
// 如果是空节点，其值为-1
func BuildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return &TreeNode{val: -1}
	}
	root := &TreeNode{val: nums[0]}
	queue := []*TreeNode{root}
	for i := 1; i < len(nums); {
		node := queue[0]
		queue = queue[1:]
		if nums[i] != -1 {
			node.left = &TreeNode{val: nums[i]}
			queue = append(queue, node.left)
		}
		i++
		if i < len(nums) && nums[i] != -1 {
			node.right = &TreeNode{val: nums[i]}
			queue = append(queue, node.right)
		}
		i++
	}
	return root
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	nums3 := make([]int, m+n)
	for i, j, k := 0, 0, 0; k < m+n; k++ {
		if i < m && j < n && nums1[i] <= nums2[j] {
			nums3[k] = nums1[i]
			i++
		} else if j < n {
			nums3[k] = nums2[j]
			j++
		} else {
			nums3[k] = nums1[i]
			i++
		}
	}
	copy(nums1, nums3)
}
