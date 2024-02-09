package main

import (
	"fmt"
	"sort"
)

type Ordered interface{
	~int | ~float64 | ~string
}

type Student struct{
	Name string
	ID int
	Age float64
}

// Group of funcitons that ensure that an OrdenedSlice can be sorted
type OrdenedSlice[T Ordered][]T // T must implement < and >

func(s OrdenedSlice[T]) Len() int {
	return len(s)
}

func(s OrdenedSlice[T]) Less(i, j int) bool {
	return s[i] < s[j]
}

func(s OrdenedSlice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
// end group for OrdenedSlice

// Group of functions that ensure that SortType can be sorted
type SortType[T any] struct {
	slice []T
	compare func(T, T) bool
}

func(s SortType[T]) Len() int {
	return len(s.slice)
}

func(s SortType[T]) Less(i, j int) bool {
	return s.compare(s.slice[i], s.slice[j])
}

func(s SortType[T]) Swap(i, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}
// end of group SortType

func PeformSort[T any](slice []T, compare func(T, T) bool) {
	sort.Sort(SortType[T]{slice, compare})
}

func addStudent[T any](students []T, student T)[]T {
	return append(students, student)
}

func main() {
	students := []string{}
	result := addStudent[string](students, "Michael")
	result = addStudent[string](result, "Jennifer")
	result = addStudent[string](result, "Elaine")
	sort.Sort(OrdenedSlice[string](result))
	fmt.Println(result)

	students1 := []int{}
	result1 := addStudent[int](students1, 78)
	result1 = addStudent[int](result1, 64)
	result1 = addStudent[int](result1, 45)
	sort.Sort(OrdenedSlice[int](result1))
	fmt.Println(result1)

	students2 := []Student{}
	result2 := addStudent[Student](students2, Student{"John", 213, 17.5})
	result2 = addStudent[Student](result2, Student{"James", 111, 18.75})
	result2 = addStudent(result2, Student{"Marsha", 110, 16.25})
	PeformSort[Student](result2, func(s1, s2 Student) bool {
		return s1.Age < s2.Age // comparing two Student values
	})
	fmt.Println(result2)
}