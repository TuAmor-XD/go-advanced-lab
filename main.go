package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

/* =======================
   PART 1: MATH FUNCTIONS
   ======================= */

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, errors.New("prime check requires number >= 2")
	}
	if n == 2 {
		return true, nil
	}
	if n%2 == 0 {
		return false, nil
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("negative exponents not supported")
	}
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result, nil
}

/* =======================
   PART 2: CLOSURES
   ======================= */

func MakeCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func MakeAccumulator(initial int) (func(int), func(int), func() int) {
	value := initial
	add := func(x int) {
		value += x
	}
	sub := func(x int) {
		value -= x
	}
	get := func() int {
		return value
	}
	return add, sub, get
}

/* =======================
   PART 3: HIGHER-ORDER
   ======================= */

func Apply(nums []int, op func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = op(v)
	}
	return result
}

func Filter(nums []int, predicate func(int) bool) []int {
	var result []int
	for _, v := range nums {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce(nums []int, initial int, op func(int, int) int) int {
	acc := initial
	for _, v := range nums {
		acc = op(acc, v)
	}
	return acc
}

func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

/* =======================
   PART 4: PROCESS
   ======================= */

func ExploreProcess() {
	fmt.Println("=== Process Information ===")
	fmt.Println("Current PID:", os.Getpid())
	fmt.Println("Parent PID:", os.Getppid())

	data := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice address: %p\n", &data)
	fmt.Printf("First element address: %p\n", &data[0])
	fmt.Println("Other processes cannot access this memory due to process isolation")
}

/* =======================
   PART 5: POINTERS
   ======================= */

func DoubleValue(x int) {
	x = x * 2
	// This will NOT modify the original variable because Go passes values by copy
}

func DoublePointer(x *int) {
	*x = *x * 2
	// This WILL modify the original variable because we use a pointer
}

func CreateOnStack() int {
	x := 10
	// This variable stays on the stack
	return x
}

func CreateOnHeap() *int {
	x := 20
	// This variable escapes to the heap
	return &x
}

func SwapValues(a, b int) (int, int) {
	return b, a
}

func SwapPointers(a, b *int) {
	*a, *b = *b, *a
}

/*
Escape Analysis Explanation:
- CreateOnHeap variable escapes because its address is returned
- Escaping means Go allocates it on the heap instead of the stack
- Stack memory is short-lived, heap memory lives longer
*/

func AnalyzeEscape() {
	CreateOnStack()
	CreateOnHeap()
}

/* =======================
   PART 6: MAIN DEMO
   ======================= */

func main() {
	ExploreProcess()

	fmt.Println("\n=== Math Operations ===")
	f, _ := Factorial(5)
	fmt.Println("Factorial(5):", f)

	p, _ := IsPrime(17)
	fmt.Println("IsPrime(17):", p)

	pw, _ := Power(2, 8)
	fmt.Println("Power(2,8):", pw)

	fmt.Println("\n=== Closures ===")
	c1 := MakeCounter(0)
	c2 := MakeCounter(100)
	fmt.Println(c1(), c1())
	fmt.Println(c2(), c1())

	fmt.Println("\n=== Higher-Order ===")
	nums := []int{1, 2, 3, 4, 5}
	squared := Apply(nums, func(x int) int { return x * x })
	fmt.Println("Squared:", squared)

	evens := Filter(nums, func(x int) bool { return x%2 == 0 })
	fmt.Println("Evens:", evens)

	sum := Reduce(nums, 0, func(a, b int) int { return a + b })
	fmt.Println("Sum:", sum)

	fmt.Println("\n=== Pointers ===")
	a, b := 5, 10
	SwapPointers(&a, &b)
	fmt.Println("After SwapPointers:", a, b)
}

