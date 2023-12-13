package main

import "fmt"

type ID interface {
	int | string
}

// 写法  [T ID, D string] == [T int | string, D string]
type UserModel[T ID, D string] struct {
	Id   T
	Name D
}

func NewUserModel[A ID, D string](id A, name D) *UserModel[A, D] {
	return &UserModel[A, D]{Id: id, Name: name}
}

func main() {
	fmt.Println(NewUserModel[int, string](10, "hello"))
	fmt.Println(NewUserModel[string, string]("10", "hello"))
}
