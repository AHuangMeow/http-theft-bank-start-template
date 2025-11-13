package main

import (
    "fmt"
    "sort"
)

func Insert(slice []int, index int, n int) []int {
    result := make([]int, 0, len(slice)+1)
    result = append(result, slice[:index]...)
    result = append(result, n)
    result = append(result, slice[index:]...)
    return result
}

func permute(nums []int) [][]int {
    if len(nums) == 1 {
        return [][]int{{nums[0]}}
    }

    permuteTail := permute(nums[1:])
    var result [][]int

    for _, arr := range permuteTail {
        for i := 0; i <= len(arr); i++ {
            result = append(result, Insert(arr, i, nums[0]))
        }
    }

    return result
}

func main() {
    var n int
    fmt.Scanf("%d", &n)

    testSlice := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&testSlice[i])
    }

    res := permute(testSlice)

    sort.Slice(res, func(i, j int) bool {
        for k := 0; k < len(res[0]); k++ {
            if res[i][k] != res[j][k] {
                return res[i][k] < res[j][k]
            }
        }
        return false
    })
    
    fmt.Println(res)
}
