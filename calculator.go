package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"os"
)

var romanToInt = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var intToRoman = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

func romanToInteger(s string) (int, error) {
	value, exists := romanToInt[s]
	if !exists {
		return 0, errors.New("Некорректное римское число")
	}
	return value, nil
}

func integerToRoman(num int) (string, error) {
	if num <= 0 {
		return "", errors.New("Результат должен быть положительным для римских чисел")
	}

	if num <= 10 {
		return intToRoman[num-1], nil
	}

	roman := ""
	if num >= 100 {
		roman += "C"
		num -= 100
	}
	if num >= 90 {
		roman += "XC"
		num -= 90
	}
	if num >= 50 {
		roman += "L"
		num -= 50
	}
	if num >= 40 {
		roman += "XL"
		num -= 40
	}
	if num >= 10 {
		roman += strings.Repeat("X", num/10)
		num %= 10
	}
	if num > 0 {
		roman += intToRoman[num-1]
	}

	return roman, nil
}

func calculate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("Ошибка: деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("Неизвестная операция")
	}
}

func isRoman(s string) bool {
	_, exists := romanToInt[s]
	return exists
}

func isArabic(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func main() {
	var input string
	fmt.Println("Введите выражение в формате: a Операция b (например: 3 + 2 или II + III)")
  
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if scanner.Scan() {
			input = scanner.Text()
		} else {
			panic("Ошибка чтения ввода")
		}

		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			panic("Неправильный ввод. Формат: a Операция b")
		}

		a := parts[0]
		op := parts[1]
		b := parts[2]

		var numA, numB int
		var err error
		if isRoman(a) && isRoman(b) {
			numA, err = romanToInteger(a)
			if err != nil {
				panic(err.Error())
			}
			numB, err = romanToInteger(b)
			if err != nil {
				panic(err.Error())
			}
		} else if isArabic(a) && isArabic(b) {
			numA, _ = strconv.Atoi(a)
			numB, _ = strconv.Atoi(b)

			if numA < 1 || numA > 10 || numB < 1 || numB > 10 {
				panic("Арабские числа должны быть в диапазоне от 1 до 10")
			}
		} else {
			panic("Числа должны быть либо арабскими, либо римскими")
		}

		result, err := calculate(numA, numB, op)
		if err != nil {
			panic(err.Error())
		}

		if isRoman(a) && isRoman(b) {
			romanResult, err := integerToRoman(result)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("Результат: %s\n", romanResult)
		} else {
			fmt.Printf("Результат: %d\n", result)
		}
	}
}
