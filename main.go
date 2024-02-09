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

func MyMap(input []int, f func(int) int) []int {
	result := make([]int, len(input))
	for index, value := range input {
		result[index] = f(value)
	}
	return result
}

func GenericMap[T1, T2 any](input []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(input))
	for index, value := range input {
		result[index] = f(value)
	}
	return result
}

func MyFilter(input []float64, f func(float64) bool)[]float64 {
	var result []float64
	for _, value := range input {
		if f(value) == true {
			result = append(result, value)
		}
	}
	return result
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

	slice := []int{1,5,2,7,4}
	result3 := MyMap(slice, func(i int) int{
		return i * i
	})
	fmt.Println(result3)

	result4 := GenericMap[int, int](slice, func(i int)int {
		return i * i
	})
	fmt.Println(result4)

	input := []float64{17.3, 11.1, 9.9, 4.3, 12.6}
	result5 := MyFilter(input, func(num float64) bool {
				return num <= 10.0
	})
	fmt.Println(result5)

}