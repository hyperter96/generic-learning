package main

import "fmt"

func RemoveDuplicate[T string | int | float64](duplicateSlice []T) []T {
	set := map[T]interface{}{}
	res := []T{}
	for _, item := range duplicateSlice {
		_, ok := set[item]
		if !ok {
			res = append(res, item)
			set[item] = nil
		}
	}
	return res
}

func main() {
	fmt.Println(RemoveDuplicate[string]([]string{"a", "c", "a"}))
	fmt.Println(RemoveDuplicate[int]([]int{1, 2, 1, 1, 1}))
}
