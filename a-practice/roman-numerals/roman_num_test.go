package main

import (
	"testing"
)

//Test v3

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
		{"4 gets converted to IV (cant repeat more than 3 times)", 4, "IV"},
		{"5 gets converted to V", 5, "V"},
		{"9 gets converted to IX", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got '%s', want '%s'", got, test.Want)
			}
		})
	}
}

//Test v2
// func TestRomanNumerals(t *testing.T) {
// 	t.Run("1 gets converted to I", func(t *testing.T) {
// 		got := ConvertToRoman(1)
// 		want := "I"

// 		if got != want {
// 			t.Errorf("got '%s', want '%s'", got, want)
// 		}
// 	})

// 	t.Run("2 gets converted to II", func(t *testing.T) {
// 		got := ConvertToRoman(2)
// 		want := "II"

// 		if got != want {
// 			t.Errorf("got '%s', want '%s'", got, want)
// 		}
// 	})
// }

//Test v1
// func TestRomanNumerals(t *testing.T) {
// 	got := ConvertToRoman(1)
// 	want := "I"

// 	if got != want {
// 		t.Errorf("got '%s', want '%s'", got, want)
// 	}
// }
