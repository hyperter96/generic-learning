package main

import "fmt"

type ModelObj interface {
	User | Product
}

type User struct {
	Uid int
}

func (this *User) SetId(id int) {
	this.Uid = id
}

type Product struct {
	Pid int
}

func (this *Product) SetId(id int) {
	this.Pid = id
}

// TrimModelObj 是一个动态类型的 Interface, 由M决定当前Interface的最终类型
type TrimModelObj[M ModelObj] interface {
	*M
	SetId(id int)
}

// TrimModelObj[Model] 由第二个参数决定当前的动态类型；
// NewModelObj[*User, User](32) 如 Model 是 User 类型, 最终 TrimModelObj == *User，所以我们需要为 Trim 传递 *User
func NewModelObj[Trim TrimModelObj[Model], Model ModelObj](id int) Trim {
	m := new(Model)
	t := Trim(m)
	fmt.Printf("%p \n", m)
	// 类型转换成指定的*Model
	t.SetId(id)
	return t
}

func main() {
	// new user model object
	user := NewModelObj[*User, User](32)
	fmt.Printf("%p \n", user)
	fmt.Printf("%T \n", user)
	fmt.Println(user.Uid)

	// new product model object
	prod := NewModelObj[*Product, Product](18)
	fmt.Printf("%p \n", prod)
	fmt.Printf("%T \n", prod)
	fmt.Println(prod.Pid)
}
