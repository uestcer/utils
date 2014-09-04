// Copyright 2014 li. All rights reserved.
// Use of this source code is governed by a MIT/X11
// license that can be found in the LICENSE file.

package collection

import (
	"testing"
)

func TestSetBasic(t *testing.T) {
	set := NewSet(1, 2, 3)
	if set.Size() != 3 || !set.Contains(1) ||
		!set.Contains(2) || !set.Contains(3) {
		t.Fatal()
	}

	if set.IsEmpty() {
		t.Fatal()
	}
}

func TestSize(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet()
	if set1.Size() != 3 || set2.Size() != 0 {
		t.Fatal()
	}
}

func TestIsEmpty(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet()
	if set1.IsEmpty() || !set2.IsEmpty() {
		t.Fatal()
	}
}

func TestContains(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet()

	if !set1.Contains(1) ||
		!set1.Contains(2) ||
		!set1.Contains(3) ||
		set1.Contains(4) {
		t.Fatal()
	}

	if set2.Contains(1) {
		t.Fatal()
	}
}

func TestToSlice(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet()

	s1 := set1.ToSlice()
	s2 := set2.ToSlice()
	if len(s1) != 3 ||
		len(s2) != 0 {
		t.Fatal()
	}
}

func TestAdd(t *testing.T) {
	set := NewSet()
	exist := set.Add(1)
	if set.Size() != 1 && set.Contains(1) && !exist {
		t.Fatal()
	}
	exist = set.Add(1)
	if exist == false {
		t.Fatal()
	}
}

func TestRemove(t *testing.T) {
	set := NewSet(1, 2, 3)
	exist := set.Remove(1)
	if set.Size() != 2 || !exist {
		t.Fatal()
	}

	exist = set.Remove(1)
	if exist {
		t.Fatal()
	}

	exist = set.Remove(4)
	if set.Size() != 2 || exist {
		t.Fatal()
	}
}

func TestClear(t *testing.T) {
	set := NewSet(1, 2, 3)
	set.Clear()
	if !set.IsEmpty() {
		t.Fatal()
	}
}

func TestUnion(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	set1.Union(set2)

	if set1.Size() != 4 ||
		!set1.Contains(1) ||
		!set1.Contains(2) ||
		!set1.Contains(3) ||
		!set1.Contains(4) {
		t.Fatal()
	}
}

func TestIntersect(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	set1.Intersect(set2)

	if set1.Size() != 2 ||
		!set1.Contains(2) || !set1.Contains(3) {
		t.Fatal()
	}
}

func TestSubtract(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	set1.Subtract(set2)

	if set1.Size() != 1 ||
		!set1.Contains(1) {
		t.Fatal()
	}
}

func TestIsSubset(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	set3 := NewSet(1, 2, 3, 4)
	if set1.IsSubset(set2) ||
		!set1.IsSubset(set1) ||
		!set1.IsSubset(set3) {
		t.Fatal()
	}
}

func TestIsEqual(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	set3 := NewSet(1, 2, 3)
	if set1.IsEqual(set2) || set2.IsEqual(set3) ||
		!set1.IsEqual(set3) {
		t.Fatal()
	}
}

func TestClone(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := set1.Clone()
	if !set1.IsEqual(set2) {
		t.Fatal()
	}
}

func TestForeach(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	sum := 0
	set1.Foreach(func(i interface{}) {
		v, _ := i.(int)
		sum += v
	})
	if sum != 6 {
		t.Fatal()
	}
}

func TestMap(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := set1.Map(func(i interface{}) interface{} {
		v, _ := i.(int)
		return v * 100
	})
	if !set2.IsEqual(NewSet(100, 200, 300)) {
		t.Fatal()
	}
}

func TestFilter(t *testing.T) {
	set1 := NewSet(1, 2, 3, 4, 5)
	set2 := set1.Filter(func(i interface{}) bool {
		v, _ := i.(int)
		return v%2 == 0
	})
	if !set2.IsEqual(NewSet(2, 4)) {
		t.Fatal()
	}
}
