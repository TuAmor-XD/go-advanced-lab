package main

import (
	"errors"
	"fmt"
)

// -------------------- Part 1 --------------------
// Factorial function
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}

// Prime check
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
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

// Power function
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

// -------------------- Part 2 --------------------
// MakeCounter
func MakeCounter(start int) func() int {
	counter := start
	return func() int {
		counter++
		return counter
	}
}

// MakeMultiplier
func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// MakeAccumulator
func MakeAccumulator(initial int) (add func(int), subtract func(int), get func() int) {
	value := initial
	add = func(x int) { value += x }
	subtract = func(x int) { value -= x }
	get = func() int { return value }
	return
}

// -------------------- Main --------------------
func main() {
	// Part 1 Tests
	fact, err := Factorial(5)
	if err != nil {
		fmt.Println("Factorial error:", err)
	} else {
		fmt.Println("Factorial 5:", fact)
	}

	prime, err := IsPrime(13)
	if err != nil {
		fmt.Println("IsPrime error:", err)
	} else {
		fmt.Println("Is 13 Prime?:", prime)
	}

	pow, err := Power(2, 3)
	if err != nil {
		fmt.Println("Power error:", err)
	} else {
		fmt.Println("2^3 =", pow)
	}

	// Part 2: MakeCounter
	counter1 := MakeCounter(0)
	fmt.Println("Counter1:", counter1()) // 1
	fmt.Println("Counter1:", counter1()) // 2

	counter2 := MakeCounter(10)
	fmt.Println("Counter2:", counter2()) // 11
	fmt.Println("Counter1:", counter1()) // 3 (independent)

	// Part 2: MakeMultiplier
	double := MakeMultiplier(2)
	triple := MakeMultiplier(3)
	fmt.Println("Double 5 =", double(5)) // 10
	fmt.Println("Triple 5 =", triple(5)) // 15

	// Part 2: MakeAccumulator
	add, subtract, get := MakeAccumulator(100)
	add(50)
	fmt.Println("Accumulator =", get()) // 150
	subtract(20)
	fmt.Println("Accumulator =", get()) // 130
}
