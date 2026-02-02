package main

import "testing"

/* ========= PART 1 TESTS ========= */

func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{"0!", 0, 1, false},
		{"1!", 1, 1, false},
		{"5!", 5, 120, false},
		{"10!", 10, 3628800, false},
		{"negative", -1, 0, true},
		{"large", 3, 6, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		input   int
		want    bool
		wantErr bool
	}{
		{2, true, false},
		{3, true, false},
		{4, false, false},
		{17, true, false},
		{1, false, true},
		{-5, false, true},
	}

	for _, tt := range tests {
		got, err := IsPrime(tt.input)
		if (err != nil) != tt.wantErr || got != tt.want {
			t.Fail()
		}
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		b, e    int
		want    int
		wantErr bool
	}{
		{2, 0, 1, false},
		{2, 3, 8, false},
		{5, 2, 25, false},
		{0, 5, 0, false},
		{2, -1, 0, true},
	}

	for _, tt := range tests {
		got, err := Power(tt.b, tt.e)
		if (err != nil) != tt.wantErr || got != tt.want {
			t.Fail()
		}
	}
}

/* ========= PART 5 TESTS ========= */

func TestSwapValues(t *testing.T) {
	a, b := SwapValues(3, 4)
	if a != 4 || b != 3 {
		t.Fail()
	}
}

func TestSwapPointers(t *testing.T) {
	a, b := 3, 4
	SwapPointers(&a, &b)
	if a != 4 || b != 3 {
		t.Fail()
	}
}
