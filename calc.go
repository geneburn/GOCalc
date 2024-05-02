package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		var firstInput, secondInput, operation string
		fmt.Println("Введите задачу, например: 2 + 2 или II + II")
		_, err := fmt.Scanf("%s %s %s", &firstInput, &operation, &secondInput)

		Operands := "0123456789IVXLC"
		Operators := "+-*/"
		switch firstInput {
		case "-I", "-II", "-III", "-IV", "-V", "-VI", "-VII", "-VIII", "-IX", "-X":
			panic("Ошибка.В римской системе нет отрицательных чисел.")
		}
		switch secondInput {
		case "-I", "-II", "-III", "-IV", "-V", "-VI", "-VII", "-VIII", "-IX", "-X":
			panic("Ошибка. В римской системе нет отрицательных чисел.")
		}

		if !strings.ContainsAny(firstInput, Operands) || !strings.ContainsAny(secondInput, Operands) {
			panic("Ошибка. Недопустимое число.")
			continue
		}
		if !strings.ContainsAny(operation, Operators) {
			fmt.Println("Ошибка. Недопустимый знак.")
			continue
		}

		if strings.ContainsAny(firstInput, "IVXLC") && !strings.ContainsAny(secondInput, "IVXLC") ||
			!strings.ContainsAny(firstInput, "IVXLC") && strings.ContainsAny(secondInput, "IVXLC") {
			panic("Ошибка.Использование двух систем счисления!")
		}

		if err != nil {
			fmt.Println("Ошибка ввода. Пожалуйста, введите задачу снова.")
			continue
		}

		firstNum, err := strconv.Atoi(firstInput)
		if err != nil {
			firstNum = convertToArabic(firstInput)
		}

		secondNum, err := strconv.Atoi(secondInput)
		if err != nil {
			secondNum = convertToArabic(secondInput)
		}

		if firstNum < -10 || firstNum > 10 || secondNum < -10 || secondNum > 10 {
			panic("Ошибка. Вводимое число должно быть не меньше -10 и не больше 10!")
		}

		fmt.Println("OK, let's GO:", firstNum, operation, secondNum)

		if operation == "+" {
			result := firstNum + secondNum
			if err != nil {
				fmt.Println("Сумма чисел =", convertToRoman(result))
			} else {
				fmt.Println("Сумма чисел =", result)
			}
		} else if operation == "-" {
			result := firstNum - secondNum
			if err != nil {
				fmt.Println("Разность чисел =", convertToRoman(result))
			} else {
				fmt.Println("Разность чисел =", result)
			}
		} else if operation == "*" {
			result := firstNum * secondNum
			if err != nil {
				fmt.Println("Произведение чисел =", convertToRoman(result))
			} else {
				fmt.Println("Произведение чисел =", result)
			}
		} else if operation == "/" {
			if secondNum == 0 {
				panic("Ошибка. На ноль делить нельзя.")
			} else {
				result := firstNum / secondNum
				if err != nil {
					fmt.Println("Частное чисел =", convertToRoman(result))
				} else {
					fmt.Println("Частное чисел =", result)
				}
			}
		}
	}
}
func convertToRoman(num int) string {
	if num == 0 || num < 0 {
		panic("В римской системе нет отрицательных чисел и нуля!")
	}

	romanNumerals := []struct {
		value   int
		numeral string
	}{
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"}, {10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	result := ""
	for _, pair := range romanNumerals {
		for num >= pair.value {
			result += pair.numeral
			num -= pair.value
		}
	}
	return result
}

func convertToArabic(roman string) int {
	romanNumerals := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
		"XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20,
		"XXI": 21, "XXII": 22, "XXIII": 23, "XXIV": 24, "XXV": 25, "XXVI": 26, "XXVII": 27, "XXVIII": 28, "XXIX": 29, "XXX": 30,
		"XXXI": 31, "XXXII": 32, "XXXIII": 33, "XXXIV": 34, "XXXV": 35, "XXXVI": 36, "XXXVII": 37, "XXXVIII": 38, "XXXIX": 39, "XL": 40,
		"XLI": 41, "XLII": 42, "XLIII": 43, "XLIV": 44, "XLV": 45, "XLVI": 46, "XLVII": 47, "XLVIII": 48, "XLIX": 49, "L": 50,
		"LI": 51, "LII": 52, "LIII": 53, "LIV": 54, "LV": 55, "LVI": 56, "LVII": 57, "LVIII": 58, "LIX": 59, "LX": 60,
		"LXI": 61, "LXII": 62, "LXIII": 63, "LXIV": 64, "LXV": 65, "LXVI": 66, "LXVII": 67, "LXVIII": 68, "LXIX": 69, "LXX": 70,
		"LXXI": 71, "LXXII": 72, "LXXIII": 73, "LXXIV": 74, "LXXV": 75, "LXXVI": 76, "LXXVII": 77, "LXXVIII": 78, "LXXIX": 79, "LXXX": 80,
		"LXXXI": 81, "LXXXII": 82, "LXXXIII": 83, "LXXXIV": 84, "LXXXV": 85, "LXXXVI": 86, "LXXXVII": 87, "LXXXVIII": 88, "LXXXIX": 89, "XC": 90,
		"XCI": 91, "XCII": 92, "XCIII": 93, "XCIV": 94, "XCV": 95, "XCVI": 96, "XCVII": 97, "XCVIII": 98, "XCIX": 99, "C": 100,
	}

	result := 0
	for i := 0; i < len(roman); i++ {
		if i+1 < len(roman) && romanNumerals[roman[i:i+2]] > 0 {
			result += romanNumerals[roman[i:i+2]]
			i++
		} else {
			result += romanNumerals[string(roman[i])]
		}
	}
	return result
}
