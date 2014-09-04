// Copyright 2014 li. All rights reserved.
// Use of this source code is governed by a MIT/X11
// license that can be found in the LICENSE file.

package collection

// Create a new set with elements.
func NewSet(elements ...interface{}) Set {
	set := &baseSet{make(map[interface{}]bool)}
	for _, element := range elements {
		set.Add(element)
	}
	return set
}

// A collection that contains no duplicate elements.
// Set is not thread safe.
type Set interface {

	// Returns the number of elements in this set (its cardinality).
	Size() int

	// Returns true if this set contains no elements.
	IsEmpty() bool

	// Returns true if this set contains the specified element.
	Contains(v interface{}) bool

	// Returns an slice containing all of the elements in this set.
	// The caller is free to modify the returned array.
	ToSlice() []interface{}

	// Adds the specified element to this set
	// Return true, if this set already contain the specified element
	Add(v interface{}) bool

	// Removes the specified element from this set
	// Return true, if this set contained the specified element
	Remove(v interface{}) bool

	// Removes all of the elements from this set.
	Clear()

	// Adds all elements in s into this set.
	Union(s Set)

	// Removes all elements not in s from this set.
	Intersect(s Set)

	// Removes all elements in s from this set.
	Subtract(s Set)

	// Returns true when all elements in this set are in s.
	IsSubset(s Set) bool

	// Returns true when two sets has the same elements.
	IsEqual(s Set) bool

	// Create a new set, and copy all the elements in this set.
	Clone() Set

	// Iterate the set elements and invoke f by every element.
	Foreach(f func(interface{}))

	// Create a new set, mapping the elements by call f.
	Map(f func(interface{}) interface{}) Set

	// Create a new set with all elements satisfied f.
	Filter(f func(interface{}) bool) Set
}

type baseSet struct {
	elements map[interface{}]bool
}

func (s *baseSet) Size() int {
	return len(s.elements)
}

func (s *baseSet) IsEmpty() bool {
	return s.Size() == 0
}

func (s *baseSet) Contains(v interface{}) bool {
	_, ok := s.elements[v]
	return ok
}

func (s *baseSet) ToSlice() []interface{} {
	values := make([]interface{}, s.Size())
	i := 0
	for k := range s.elements {
		values[i] = k
		i++
	}
	return values
}

func (s *baseSet) Add(v interface{}) bool {
	_, ok := s.elements[v]
	s.elements[v] = true
	return ok
}

func (s *baseSet) Remove(v interface{}) bool {
	_, ok := s.elements[v]
	if ok {
		delete(s.elements, v)
	}
	return ok
}

func (s *baseSet) Clear() {
	s.elements = make(map[interface{}]bool)
}

func (s0 *baseSet) Union(s1 Set) {
	if s1 == nil {
		return
	}
	s1.Foreach(func(i interface{}) {
		s0.Add(i)
	})
}

func (s *baseSet) Intersect(s1 Set) {
	if s1 == nil {
		return
	}
	for k := range s.elements {
		if !s1.Contains(k) {
			delete(s.elements, k)
		}
	}
}

func (s *baseSet) Subtract(s1 Set) {
	if s1 == nil {
		return
	}

	s1.Foreach(func(i interface{}) {
		s.Remove(i)
	})
}

func (s *baseSet) IsSubset(s1 Set) bool {
	if s1 == nil || s.Size() > s1.Size() {
		return false
	}

	for k, _ := range s.elements {
		if !s1.Contains(k) {
			return false
		}
	}
	return true
}

func (s0 *baseSet) IsEqual(s1 Set) bool {
	if s1 == nil || s0.Size() != s1.Size() {
		return false
	}

	for k, _ := range s0.elements {
		if !s1.Contains(k) {
			return false
		}
	}
	return true
}

func (s *baseSet) Clone() Set {
	elements := make(map[interface{}]bool)
	for k := range s.elements {
		elements[k] = true
	}
	return &baseSet{elements}
}

func (s *baseSet) Foreach(f func(interface{})) {
	for k, _ := range s.elements {
		f(k)
	}
}

func (s *baseSet) Map(f func(interface{}) interface{}) Set {
	result := NewSet()
	for k, _ := range s.elements {
		result.Add(f(k))
	}
	return result
}

func (s *baseSet) Filter(f func(interface{}) bool) Set {
	result := NewSet()
	for k, _ := range s.elements {
		if f(k) {
			result.Add(k)
		}
	}
	return result
}
