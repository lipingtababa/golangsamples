package main

import "fmt"
import "strings"
import "errors"
import "strconv"

//Doesn't support **, ^, ! or square
//All numbers are treated as float64
var SUPPORT_OPERATORS = []string{"+", "-", "*", "/"}

func main() {
	//TODO should use unittest, if there is one
	var expressions []string
	expressions = []string{"3 2 1 - /"}
	expressions = []string{"3 + 1"}
	expressions = []string{"1 2 +", "1 2 3 * +", "3 2 1 / *", "3 2 1", "3 + -", " 1 2 + 3 4 - * ", "3 2 + 4 ^", "", "3 2 2 - /", "3 2 **", "+ - * /", "3 + 1"}

	for _, expr := range expressions {
		result, err := calculate(expr)
		if err == nil {
			fmt.Println("expr [", expr, "] equals ", result)
		} else {
			fmt.Println("expr [", expr, "] is invalid because ", err)
		}
	}
}

func calculate(expression string) (float64, error) {
	var tokens = strings.Fields(expression)
	operands := []float64{}
	var operand1, operand2 float64
	var err error = nil
	for _, token := range tokens {
		if isSupportedOperator(token) {
			//TODO Assume that there would be enough operands
			operand2, operands, err = popFloat(operands)
			if err != nil {
				return -1, errors.New(fmt.Sprintf("Not enough operands for operator %s", token))
			}
			operand1, operands, err = popFloat(operands)
			if err != nil {
				return -1, errors.New(fmt.Sprintf("Not enough operands for operator %s and operand %f ", token, operand2))
			}
			fmt.Println("Calculate ", operand1, token, operand2)
			var result float64
			switch token {
			case "+":
				result = operand1 + operand2
			case "-":
				result = operand1 - operand2
			case "*":
				result = operand1 * operand2
			case "/":
				//TODO assume there would be no divisionbyzero error
				if operand2 == 0 {
					return -1, errors.New(fmt.Sprint(operand1, " is being divided by Zero"))
				}
				result = operand1 / operand2
			}
			operands = prependFloat(operands, result)
		} else {
			operand, err := strconv.ParseFloat(token, 32)
			if err == nil {
				operands = prependFloat(operands, operand)
				fmt.Println("Append ", operand)
				//fmt.Println("Operends ", operands)
			} else {
				return -1, errors.New(fmt.Sprintf("[%s] is Not a Valid operand", token))
			}
		}
	}
	if len(operands) >= 1 {
		//Return the latest operand
		return operands[0], nil
	} else {
		//Empty expression
		return 0, nil
	}
}

func popFloat(slice []float64) (header float64, output []float64, err error) {
	if len(slice) < 1 {
		return 0, slice, errors.New("No value to be popped")
	}
	header = slice[0]
	//The tmp must be of enough capacity, otherwise the copy will fail
	tmp := make([]float64, len(slice)-1)
	copy(tmp, slice[1:])
	return header, tmp, nil
}

func prependFloat(slice []float64, header float64) (output []float64) {
	slice = append(slice, 0)
	copy(slice[1:], slice)
	slice[0] = header
	return slice
}

func isSupportedOperator(operator string) bool {
	for _, op := range SUPPORT_OPERATORS {
		if operator == op {
			return true
		}
	}
	return false
}
