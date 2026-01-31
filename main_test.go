package main

import (
	"testing"
)

// -------------------- Part 1: Factorial --------------------
func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{name: "factorial of 0", input: 0, want: 1, wantErr: false},
		{name: "factorial of 1", input: 1, want: 1, wantErr: false},
		{name: "factorial of 3", input: 3, want: 6, wantErr: false},
		{name: "factorial of 5", input: 5, want: 120, wantErr: false},
		{name: "factorial of 6", input: 6, want: 720, wantErr: false},
		{name: "factorial of -1 (error)", input: -1, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factorial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -------------------- Part 2: MakeCounter --------------------
func TestMakeCounter(t *testing.T) {
	tests := []struct {
		name   string
		start  int
		calls  int
		expect []int
	}{
		{name: "counter from 0, 3 calls", start: 0, calls: 3, expect: []int{1, 2, 3}},
		{name: "counter from 10, 2 calls", start: 10, calls: 2, expect: []int{11, 12}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := MakeCounter(tt.start)
			for i, want := range tt.expect {
				got := counter()
				if got != want {
					t.Errorf("Call %d: got %d, want %d", i+1, got, want)
				}
			}
		})
	}

	// Test independence
	counterA := MakeCounter(0)
	counterB := MakeCounter(100)
	counterA() // 1
	counterB() // 101
	counterA() // 2
	if counterA() != 3 {
		t.Errorf("CounterA did not increment independently")
	}
	if counterB() != 102 {
		t.Errorf("CounterB did not increment independently")
	}
}

// -------------------- Part 2: MakeMultiplier --------------------
func TestMakeMultiplier(t *testing.T) {
	tests := []struct {
		name   string
		factor int
		input  int
		want   int
	}{
		{name: "double 5", factor: 2, input: 5, want: 10},
		{name: "triple 7", factor: 3, input: 7, want: 21},
		{name: "quadruple 4", factor: 4, input: 4, want: 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mult := MakeMultiplier(tt.factor)
			got := mult(tt.input)
			if got != tt.want {
				t.Errorf("MakeMultiplier(%d)(%d) = %d; want %d", tt.factor, tt.input, got, tt.want)
			}
		})
	}
}

// -------------------- Part 2: MakeAccumulator --------------------
func TestMakeAccumulator(t *testing.T) {
	add, subtract, get := MakeAccumulator(100)

	add(50)
	if got := get(); got != 150 {
		t.Errorf("After add(50), got %d; want 150", got)
	}

	subtract(20)
	if got := get(); got != 130 {
		t.Errorf("After subtract(20), got %d; want 130", got)
	}

	add(10)
	subtract(15)
	if got := get(); got != 125 {
		t.Errorf("After add(10) and subtract(15), got %d; want 125", got)
	}
}
