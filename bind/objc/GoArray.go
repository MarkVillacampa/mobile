package main

/*
#cgo CFLAGS: -x objective-c -fobjc-arc -fmodules -fblocks -Werror
#cgo LDFLAGS: -framework Foundation

#include <stdint.h>
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"reflect"

	_seq "golang.org/x/mobile/bind/seq"
)

////export go_objectAtIndex
//func go_objectAtIndex(refnum C.int32_t, index C.int) C.int32_t {
//  ref := _seq.FromRefNum(int32(refnum))
//  v := ref.Get()
//  o := reflect.ValueOf(v).Elem().Index(int(index))
//  fmt.Printf("(%v, %T)\n", o, o)
//  id := C.int32_t(_seq.ToRefNum(o))
//  fmt.Printf("(%v, %T)\n", id, id)
//  return id
//}
//// func go_objectAtIndex(refnum C.int32_t, index C.int) C.int {
////   ref := _seq.FromRefNum(int32(refnum))
////   v := *ref.Get().(*[]int)
////   return C.int(v[int(index)])
//// }

////export go_count
//func go_count(refnum C.int32_t) C.int {
//  ref := _seq.FromRefNum(int32(refnum))
//  v := *ref.Get().(*[]int)
//  return C.int(len(v))
//}

////export go_insertObjectAtIndex
//func go_insertObjectAtIndex(refnum C.int32_t, obj C.int, index C.int) {
//  ref := _seq.FromRefNum(int32(refnum))
//  v := *ref.Get().(*[]int)
//  v[int(index)] = int(obj)
//}

////export go_removeObjectAtIndex
//func go_removeObjectAtIndex(refnum C.int32_t, index C.int) {
//  ref := _seq.FromRefNum(int32(refnum))
//  v := *ref.Get().(*[]int)
//  v = append(v[:index], v[index+1:]...)
//}

////export go_addObject
//func go_addObject(refnum C.int32_t, obj C.int) {
//  ref := _seq.FromRefNum(int32(refnum))
//  v := *ref.Get().(*[]int)
//  v = append(v, int(obj))
//}

////export go_removeLastObject
//func go_removeLastObject(refnum C.int32_t) {
//  ref := _seq.FromRefNum(int32(refnum))
//  v := *ref.Get().(*[]int)
//  v = v[:len(v)-1]
//}

////export go_replaceObjectAtIndexWithObject
//func go_replaceObjectAtIndexWithObject(refnum C.int32_t, index C.int, obj C.int) {
//  ref := _seq.FromRefNum(int32(refnum))
//  v := *ref.Get().(*[]int)
//  v[int(index)] = int(obj)
//}

////export go_new_slice
//func go_new_slice() C.int32_t {
//  a := []int{99, 98, 97, 96, 95, 1,2,3}
//  return C.int32_t(_seq.ToRefNum(&a))
//  // return C.int32_t(_seq.ToRefNum(new([]int)))
//}

//export go_objectAtIndex
func go_objectAtIndex(refnum C.int32_t, index C.int) C.int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get()
	o := reflect.ValueOf(v).Elem()
	arr := reflect.ValueOf(o.Interface())
	obj := arr.Index(int(index))
	fmt.Printf("go_objectAtIndex(%d): (%v, %T)\n", int(index), obj, obj)
	return C.int32_t(_seq.ToRefNum(obj.Interface()))
}

//export go_count
func go_count(refnum C.int32_t) C.int {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get()
	l := reflect.ValueOf(v).Elem()
	arr := reflect.ValueOf(l.Interface())
	return C.int(arr.Len())
}

//export go_insertObjectAtIndex
func go_insertObjectAtIndex(refnum C.int32_t, objref C.int32_t, index C.int) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get()
	o := reflect.ValueOf(v).Elem()
	arr := reflect.ValueOf(o.Interface())
	obj := _seq.FromRefNum(int32(objref)).Get()
	arr.Index(int(index)).Set(reflect.ValueOf(obj))
	// v[int(index)] = obj
}

//export go_removeObjectAtIndex
func go_removeObjectAtIndex(refnum C.int32_t, index C.int) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get()
	optr := reflect.ValueOf(v)
	o := optr.Elem()
	arr := reflect.ValueOf(o.Interface())
	first := arr.Slice(0, int(index))
	second := arr.Slice(int(index)+1, arr.Len())
	newslice := reflect.AppendSlice(first, second)
	o.Set(newslice)
	// v = append(v[:index], v[index+1:]...)
}

//export go_addObject
func go_addObject(refnum C.int32_t, objref C.int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get()
	o := reflect.ValueOf(v).Elem()
	arr := reflect.ValueOf(o.Interface())
	fmt.Printf("go_addObject: (%v, %T)\n", v, v)
	fmt.Printf("go_addObject: (%v, %T)\n", o, o)
	obj := _seq.FromRefNum(int32(objref)).Get()
	o.Set(reflect.Append(arr, reflect.ValueOf(obj)))
	// v = append(v, int(obj))
}

//export go_removeLastObject
func go_removeLastObject(refnum C.int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get()
	o := reflect.ValueOf(v).Elem()
	arr := reflect.ValueOf(o.Interface())
	o.Set(arr.Slice(0, arr.Len()-1))
	// v = v[:len(v)-1]
}

//export go_replaceObjectAtIndexWithObject
func go_replaceObjectAtIndexWithObject(refnum C.int32_t, index C.int, objref C.int32_t) {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get()
	o := reflect.ValueOf(v).Elem()
	arr := reflect.ValueOf(o.Interface())
	obj := _seq.FromRefNum(int32(objref)).Get()
	arr.Index(int(index)).Set(reflect.ValueOf(obj))
	// v[int(index)] = int(obj)
}

//export go_new_slice
func go_new_slice(refnum C.int32_t) C.int32_t {
	ref := _seq.FromRefNum(int32(refnum))
	v := ref.Get()
	t := reflect.TypeOf(v)
	a := reflect.MakeSlice(reflect.SliceOf(t), 0, 0)
	fmt.Printf("go_new_slice: (%v, %T)\n", a, a)
	i := a.Interface()
	fmt.Printf("go_new_slice: (%v, %T)\n", i, i)
	return C.int32_t(_seq.ToRefNum(&i))
}
