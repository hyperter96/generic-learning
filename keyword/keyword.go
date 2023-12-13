package main

import "fmt"

func DefaultKeyWordParams[D any](defVal D, params ...D) D {
	if len(params) == 0 {
		return defVal
	}
	return params[0]
}

func test(category ...string) {
	// 不填写则返回默认值
	realCategory := DefaultKeyWordParams[string]("AGroup", category...)
	fmt.Println(realCategory)
}

func main() {
	test()
}
