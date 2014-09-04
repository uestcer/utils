// Copyright 2014 li. All rights reserved.
// Use of this source code is governed by a MIT/X11
// license that can be found in the LICENSE file.

package slice

import (
	"container/list"
	"reflect"
	"testing"
)

func TestAsList(t *testing.T) {
	l1 := AsList(1, 2, "3", "Hello, GoLang")

	l2 := list.New()
	l2.PushBack(1)
	l2.PushBack(2)
	l2.PushBack("3")
	l2.PushBack("Hello, GoLang")

	if !reflect.DeepEqual(l1, l2) {
		t.Fatal()
	}
}

func TestToList(t *testing.T) {
	s := []int{1, 2, 3, 4}

	l1 := list.New()
	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)
	l1.PushBack(4)

	l2 := ToList(s)
	l3 := ToList(s)

	if !reflect.DeepEqual(l1, l2) || !reflect.DeepEqual(l2, l3) {
		t.Fatal()
	}
}

func TestToSlice(t *testing.T) {
	l1 := list.New()
	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)
	l1.PushBack(4)

	if !reflect.DeepEqual(ToSlice(l1), []interface{}{1, 2, 3, 4}) {
		t.Fatal()
	}
}

func TestForeach(t *testing.T) {
	sum := 0
	Foreach([]int{1, 2, 3, 4}, func(i int) { sum += i })
	if sum != 10 {
		t.Fatal()
	}
}

func TestMap(t *testing.T) {
	r := Map([]int{1, 2, 3, 4}, func(i int) int { return i * 100 })
	if !reflect.DeepEqual(r, []interface{}{100, 200, 300, 400}) {
		t.Fatal()
	}
}

func TestExist(t *testing.T) {
	r1 := Exist([]int{1, 2, 3, 4}, func(i int) bool { return i%3 == 0 })
	r2 := Exist([]int{1, 2, 3, 4}, func(i int) bool { return i%5 == 0 })
	if r1 == false {
		t.Fatal()
	}
	if r2 == true {
		t.Fatal()
	}
}

func TestFilter(t *testing.T) {
	rs := Filter([]int{1, 2, 3, 4}, func(i int) bool { return i%2 == 0 })
	if !reflect.DeepEqual([]interface{}{2, 4}, rs) {
		t.Fatal()
	}
}

func TestFind(t *testing.T) {
	ok1, r1 := Find([]int{1, 2, 3, 4, 6}, func(i int) bool { return i%3 == 0 })
	ok2, _ := Find([]int{1, 2, 3, 4}, func(i int) bool { return i%5 == 0 })

	n1, _ := r1.(int)
	if ok1 != true || n1 != 3 {
		t.Fatal()
	}
	if ok2 != false {
		t.Fatal()
	}
}

func TestFindLast(t *testing.T) {
	ok1, r1 := FindLast([]int{1, 2, 3, 4, 6}, func(i int) bool { return i%3 == 0 })
	ok2, _ := FindLast([]int{1, 2, 3, 4}, func(i int) bool { return i%5 == 0 })

	n1, _ := r1.(int)
	if ok1 != true || int(n1) != 6 {
		t.Fatal()
	}
	if ok2 != false {
		t.Fatal()
	}
}
