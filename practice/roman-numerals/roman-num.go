package main

import "strings"

//Solution v5
type RomanNumeral struct {
	Value  int
	Symbol string
}

var RomanNumerals = []RomanNumeral{
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range RomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

// Solution v4
// func ConvertToRoman(arabic int) string {
// 	var result strings.Builder

// 	for arabic > 0 {
// 			switch {
// 			case arabic > 9:
// 				result.WriteString("X")
// 				arabic -= 10
// 			case arabic > 8:
// 				result.WriteString("IX")
// 				arabic -= 9
// 			case arabic > 4:
// 				result.WriteString("V")
// 				arabic -= 5
// 			case arabic > 3:
// 				result.WriteString("IV")
// 				arabic -=4
// 			default:
// 				result.WriteString("I")
// 				arabic--
// 			}
// 	}
// 	return result.String()
// }

// Solution v3
// func ConvertToRoman(arabic int) string {
// 	var result strings.Builder

// 	for i:=arabic; i>0; i-- {
// 		if i == 5 {
// 			result.WriteString("V")
// 			break
// 		}

// 		if i == 4 {
// 			result.WriteString("IV")
// 			break
// 		}
// 		result.WriteString("I")
// 	}
// 	return result.String()
// }

//Solution v2
// func ConvertToRoman(arabic int) string {

// 	if arabic == 4 {
// 		return "IV"
// 	}

// 	var result strings.Builder

// 	for i:=0; i<arabic; i++ {
// 		result.WriteString("I")
// 	}

// 	return result.String()
// }

//Solution v1
// func ConvertToRoman(arabic int) string {
// 	if arabic == 3 {
// 		return "III"
// 	}
// 	if arabic == 2 {
// 		return "II"
// 	}
// 	return "I"
// }
