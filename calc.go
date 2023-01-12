package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	listOfNumbers := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	operations := "+-*/"
	for continueLoop := true; continueLoop; {
		fmt.Println("Введите выражение в формате А + В, где А и В - числа от 1 до 10 " +
			"(так же принимаются в виде римских от I до X), а вместо + любая операция из + - * и /.")
		fmt.Print("->")
		expression, _ := reader.ReadString('\n')
		expressionSlice := strings.Split(strings.TrimSpace(expression), " ")
		if len(expressionSlice) != 3 {
			fmt.Println("Не верно введен формат выражения")
			break
		}

		var isFirstVariableRoman bool
		if v, err := strconv.Atoi(expressionSlice[0]); err == nil {
			if v > 0 && v <= 10 {
				isFirstVariableRoman = false
			} else {
				fmt.Println("Не верно введен формат выражения")
				break
			}
		} else if _, ok := listOfNumbers[expressionSlice[0]]; ok == true {
			isFirstVariableRoman = true
		} else {
			fmt.Println("Не верно введен формат выражения")
			break
		}
		var isSecondVariableRoman bool
		if v, err := strconv.Atoi(expressionSlice[2]); err == nil {
			if v > 0 && v <= 10 {
				isSecondVariableRoman = false
			} else {
				fmt.Println("Не верно введен формат выражения")
				break
			}
		} else if _, ok := listOfNumbers[expressionSlice[2]]; ok == true {
			isSecondVariableRoman = true
		} else {
			fmt.Println("Не верно введен формат выражения")
			break
		}

		if strings.Contains(operations, expressionSlice[1]) {
			if (isFirstVariableRoman && isSecondVariableRoman) || (!isFirstVariableRoman && !isSecondVariableRoman) {
				var a, b int
				if isFirstVariableRoman {
					a, b = listOfNumbers[expressionSlice[0]], listOfNumbers[expressionSlice[2]]
				} else {
					a, _ = strconv.Atoi(expressionSlice[0])
					b, _ = strconv.Atoi(expressionSlice[2])
				}

				switch {
				case expressionSlice[1] == "+":
					if isFirstVariableRoman {
						fmt.Printf("%s + %s = %s\n", expressionSlice[0], expressionSlice[2], intToRoman(a+b))
					} else {
						fmt.Printf("%d + %d = %d\n", a, b, a+b)
					}
				case expressionSlice[1] == "-":
					if isFirstVariableRoman {
						if a-b < 1 {
							fmt.Println("Результат 0 или отрицателен")
							continueLoop = false
							break
						}
						fmt.Printf("%s - %s = %s\n", expressionSlice[0], expressionSlice[2], intToRoman(a-b))
					} else {
						fmt.Printf("%d - %d = %d\n", a, b, a-b)
					}
				case expressionSlice[1] == "*":
					if isFirstVariableRoman {
						fmt.Printf("%s * %s = %s\n", expressionSlice[0], expressionSlice[2], intToRoman(a*b))
					} else {
						fmt.Printf("%d * %d = %d\n", a, b, a*b)
					}
				case expressionSlice[1] == "/":
					if isFirstVariableRoman {
						if a/b < 1 {
							fmt.Println("Результат меньше 1")
							continueLoop = false
							break
						}
						fmt.Printf("%s / %s = %s\n", expressionSlice[0], expressionSlice[2], intToRoman(a/b))
					} else {
						fmt.Printf("%d / %d = %d\n", a, b, a/b)
					}
				}
			} else {
				fmt.Println("Числа в разной системе счисления")
				break
			}
		} else {
			fmt.Println("Неверная математическая операция")
			break
		}
	}
}

func intToRoman(intResult int) string {
	roman := ""
	ints := []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	romans := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	for i := len(ints) - 1; intResult > 0; i-- {
		for ints[i] <= intResult {
			roman += romans[i]
			intResult -= ints[i]
		}
	}
	return roman
}
