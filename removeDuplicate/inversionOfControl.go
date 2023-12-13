package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func NewStudent(name string, age int) *Student {
	return &Student{Name: name, Age: age}
}

func DefaultFilter(item interface{}) (uniqueKey interface{}) {
	return item.(*Student).Name
}

func RemoveDuplicateWithFilter[T comparable](compareSlice []T, filterFunc func(item interface{}) (key interface{})) []T {
	set := map[interface{}]interface{}{}
	res := []T{}
	for _, item := range compareSlice {
		i := filterFunc(item)
		_, ok := set[i]
		if !ok {
			res = append(res, item)
			set[i] = nil
		}
	}
	return res
}

func main() {
	s := []*Student{
		NewStudent("a", 1),
		NewStudent("a", 1),
		NewStudent("b", 2),
		NewStudent("b", 2),
	}
	l := RemoveDuplicateWithFilter[*Student](s, DefaultFilter)
	for _, i := range l {
		fmt.Println(i.Name, i.Age)
	}
}
