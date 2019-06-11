package main

import (
	"testing"
)

// func TestAdd(t *testing.T) {
// 	got := Add(2, 2)
// 	want := 4
// 	if got != want {
// 		t.Errorf("got %d, want %d", got, want)
// 	}
// }

// func TestAddValues(t *testing.T) {
// 	got := Add(3, 3)
// 	want := 6
// 	if got != want {
// 		t.Errorf("got %d, want %d", got, want)
// 	}
// }

// func Add(a, b int) int {
// 	return a + b
// }

// func ExampleAdd() {
// 	sum := Add(1, 5)
// 	fmt.Println(sum)
// 	// Output: 6
// }

// func TestAddMoreThan2Numbers(t *testing.T) {
// 	got := Add(1, 2, 3)
// 	want := 6

// 	if got != want {
// 		t.Errorf("got %d, want %d", got, want)
// 	}
// }

func Add(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func TestAdd(t *testing.T) {
	cases := []struct {
		Description string
		Numbers     []int
		Want        int
	}{
		{Description: "Add two numbers", Numbers: []int{1, 2}, Want: 3},
		{Description: "Add three numbers", Numbers: []int{1, 2, 3}, Want: 6},
		//{Description: "Add all numbers apart from first one ", Numbers: []int{1, 2, 3}, Want: 5},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := Add(test.Numbers...)
			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}

		})
	}
}
