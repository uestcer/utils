// Copyright 2014 li. All rights reserved.
// Use of this source code is governed by a MIT/X11
// license that can be found in the LICENSE file.

// Useful functions for handle slice.
// NOTE: function will panic if the argument type is not
// correct at runtime.
package slice

import (
	"container/list"
	"reflect"
)

// New a list, and append the elements to list in order.
// Example: slice.AsList(1, 2, "3", "Hello, GoLang")
func AsList(elements ...interface{}) *list.List {
	l := list.New()
	for _, element := range elements {
		l.PushBack(element)
	}
	return l
}

// Convert the slice elements to list.
// Panic if the argument type is not slice or slice pointer.
// Example: slice.ToList(&[]int{1, 2, 3, 5}), slice.ToList([]int{100, 2, 3, 5})
func ToList(s interface{}) *list.List {
	v := reflectSlice(s)
	l := list.New()
	for i := 0; i < v.Len(); i++ {
		l.PushBack(v.Index(i).Interface())
	}
	return l
}

// Convert the list elements to slice.
// Return empty slice ,if list if nil or empty;
func ToSlice(l *list.List) []interface{} {
	if l == nil {
		return make([]interface{}, 0)
	}

	result := make([]interface{}, l.Len())
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		result[i] = e.Value
		i++
	}
	return result
}

// Traverse the slice, call function f by element in order.
// NOTE: Panic if i is not slice or slice pointer, f type is not func or func pointer.
func Foreach(i interface{}, f interface{}) {
	v1 := reflectSlice(i)
	v2 := reflectFunc(f)

	for i := 0; i < v1.Len(); i++ {
		v2.Call([]reflect.Value{v1.Index(i)})
	}
}

// Map the slice to another slice, convert element by function f in order.
// NOTE: Panic if i is not slice or slice pointer, f type is not func or func pointer.
func Map(i interface{}, f interface{}) []interface{} {
	v1 := reflectSlice(i)
	v2 := reflectFunc(f)

	result := make([]interface{}, v1.Len())
	for i := 0; i < v1.Len(); i++ {
		result[i] = v2.Call([]reflect.Value{v1.Index(i)})[0].Interface()
	}
	return result
}

// Check if the slice has element satisfy function f.
// NOTE: Panic if i is not slice or slice pointer, f type is not func or func pointer.
// Return true if slice has at least such one element, Otherwise false.
func Exist(i interface{}, f interface{}) bool {
	v1 := reflectSlice(i)
	v2 := reflectFunc(f)

	for i := 0; i < v1.Len(); i++ {
		if v2.Call([]reflect.Value{v1.Index(i)})[0].Bool() {
			return true
		}
	}
	return false
}

// Filter element satisfy function f, then return a new slice.
// NOTE: Panic if i is not slice or slice pointer, f type is not func or func pointer.
// If no element satisfied, return an empty slice.
func Filter(i interface{}, f interface{}) []interface{} {
	v1 := reflectSlice(i)
	v2 := reflectFunc(f)

	result := make([]interface{}, 0)
	for i := 0; i < v1.Len(); i++ {
		e := v1.Index(i)
		if v2.Call([]reflect.Value{e})[0].Bool() {
			result = append(result, e.Interface())
		}
	}
	return result
}

// Get first element index satisfy function f
// NOTE: Panic if i is not slice or slice pointer, f type is not func or func pointer.
// Return -1, if no element satisfy.
func Index(i interface{}, f interface{}) int {
	v1 := reflectSlice(i)
	v2 := reflectFunc(f)

	for i := 0; i < v1.Len(); i++ {
		e := v1.Index(i)
		if v2.Call([]reflect.Value{e})[0].Bool() {
			return i
		}
	}
	return -1
}

// Get first element index satisfy function f in reverse order.
// NOTE: Panic if i is not slice or slice pointer, f type is not func or func pointer.
// Return -1, if no element satisfy.
func IndexLast(i interface{}, f interface{}) int {
	v1 := reflectSlice(i)
	v2 := reflectFunc(f)

	for i := v1.Len() - 1; i > 0; i-- {
		e := v1.Index(i)
		if v2.Call([]reflect.Value{e})[0].Bool() {
			return i
		}
	}
	return -1
}

// Find first element satisfy function f
// NOTE: Panic if i is not slice or slice pointer, f type is not func or func pointer.
func Find(i interface{}, f interface{}) (bool, interface{}) {
	v1 := reflectSlice(i)
	v2 := reflectFunc(f)

	for i := 0; i < v1.Len(); i++ {
		e := v1.Index(i)
		if v2.Call([]reflect.Value{e})[0].Bool() {
			return true, e.Interface()
		}
	}
	return false, nil
}

// Find first element satisfy function f in reverse order.
// NOTE: Panic if i is not slice or slice pointer, f type is not func or func pointer.
func FindLast(i interface{}, f interface{}) (bool, interface{}) {
	v1 := reflectSlice(i)
	v2 := reflectFunc(f)

	for i := v1.Len() - 1; i > 0; i-- {
		e := v1.Index(i)
		if v2.Call([]reflect.Value{e})[0].Bool() {
			return true, e.Interface()
		}
	}
	return false, nil
}

// Reflect i to reflect.Value, Elem() if value is PTR.
// NOTE: Panic if the argument type is not func or func pointer.
func reflectFunc(f interface{}) reflect.Value {
	v := reflect.ValueOf(f)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Func {
		panic("utils/slice: argument type is not func, " + v.Kind().String() + ".")
	}
	return v
}

// Reflect i to reflect.Value, Elem() if value is PTR.
// NOTE: Panic if the argument type is not slice or slice pointer.
func reflectSlice(i interface{}) reflect.Value {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Slice {
		panic("utils/slice: argument type is not slice, " + v.Kind().String() + ".")
	}
	return v
}
