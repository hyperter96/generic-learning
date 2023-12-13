package main

import (
	"fmt"
	"math"
)

type KeepItem bool

// 若需要保留的item 则返回true 即可
type FilterFunc func(item interface{}) KeepItem

type PageList[T any] struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Size  int `json:"size"`
	List  []T `json:"list"`
}

type Pager[T any] struct {
	limit   int
	offset  int
	total   int
	pageCnt int
	list    []T
}

func NewPager[T any](list []T) *Pager[T] {
	return &Pager[T]{
		limit:  10,
		offset: 1,
		total:  len(list),
		list:   list,
	}
}

func (this *Pager[T]) Filter(filterFn FilterFunc) *Pager[T] {
	tmpList := []T{}
	for _, item := range this.list {
		if filterFn(&item) {
			tmpList = append(tmpList, item)
		}
	}
	this.list = tmpList
	this.total = len(tmpList)
	return this
}

func (this *Pager[T]) Offset(c int) *Pager[T] {
	this.offset = c
	return this
}

func (this *Pager[T]) Limit(c int) *Pager[T] {
	this.limit = c
	return this
}

func (this *Pager[T]) List() []T {
	// 页码
	if this.offset <= 0 {
		this.offset = 1
	}
	// size
	if this.limit > this.total {
		this.limit = this.total
	}
	// 总页数
	this.pageCnt = int(math.Ceil(float64(this.total) / float64(this.limit)))
	if this.offset > this.pageCnt {
		return []T{}
	}
	startIdx := (this.offset - 1) * this.limit
	endIdx := startIdx + this.limit

	if endIdx > this.total {
		endIdx = this.total
	}

	return this.list[startIdx:endIdx]
}

func (this *Pager[T]) Output() *PageList[T] {

	return &PageList[T]{
		Total: this.total,
		Page:  this.offset,
		Size:  this.limit,
		List:  this.list,
	}
}

// test
func main() {
	page := NewPager[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	list := page.Offset(1).Limit(3).Filter(func(item interface{}) KeepItem {
		if *item.(*int)%2 == 1 {
			return true
		}
		return false
	}).List()
	fmt.Println(list)
}
