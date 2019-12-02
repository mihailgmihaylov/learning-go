package main

import (
	"fmt"
)

// func printAsManyTimes(s, sep string, n int) string {
// 	if n == 1 {
// 		return s
// 	}
// 	return s + sep + printAsManyTimes(s, sep, n-1)
// }

// func Repeater(s, sep string) func(int) string {
// 	return func(n int) string {
// 		return printAsManyTimes(s, sep, n)
// 	}
// }

func Repeater(s, sep string) func(int) string {
	return func(n int) string {
		var result string
		for i := 1; i <= n-1; i++ {
			result = result + s + sep
		}
		fmt.Println(result + s)
		return result + s
	}
}

func Generator(gen func(int) int, initial int) func() int {
	return func() int {
		product := initial
		initial = gen(initial)
		fmt.Println(product)
		return product
	}
}

func MapReducer(mapper func(int) int, reducer func(int, int) int, initial int) func(...int) int {
	return func(args ...int) int {
		var result int
		for _, i := range args {
			temp := initial
			initial = reducer(mapper(i), initial)
			result = reducer(mapper(i), temp)
		}
		fmt.Println(result)
		return result
	}
}

func main() {
	fmt.Println("1 question (currying):")
	Repeater("foo", ":")(3)

	fmt.Println("2 question (closures):")
	counter := Generator(
		func(v int) int { return v + 1 },
		0,
	)
	power := Generator(
		func(v int) int { return v * v },
		2,
	)
	counter()
	counter()
	power()
	power()
	counter()
	power()
	counter()
	power()

	fmt.Println("3 question(arbitrary number of variables):")
	powerSum := MapReducer(
		func(v int) int { return v * v },
		func(a, v int) int { return a + v },
		0,
	)

	powerSum(1, 2, 3, 4)
}
