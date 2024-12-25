// 3. Write a function to solve the "Two Sum" problem: https://leetcode.com/problems/two-sum/ using a map.

package main

import "fmt"

func Exercise_3(nums []int, target int) []int {
    m := make(map[int]int)
    for idx, num := range nums {
        if val, exists := m[num]; exists {
            fmt.Println(idx, val)
            return []int{idx, val}
        } else {
            m[target - num] = idx
        }
    }
    return nil
}