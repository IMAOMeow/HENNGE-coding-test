package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	numTestcases, err := readSingleInteger(reader)
	if err != nil {
		log.Fatal(err)
	}

	result, err := calculateTestcasesSquareSum(reader, numTestcases, []int{})
	if err != nil {
		log.Fatal(err)
	}

	printSquareSum(result)
}

func printSquareSum(input []int) {
	if len(input) == 0 {
		return
	}

	fmt.Println(input[0])
	printSquareSum(input[1:])
}

// readSingleInteger reads single integer from standard input.
func readSingleInteger(reader *bufio.Reader) (int, error) {
	numStrTestCases, err := reader.ReadString('\n')
	if err != nil {
		return -1, err
	}

	input := strings.TrimSpace(numStrTestCases)
	numTestCases, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return -1, err
	}

	return int(numTestCases), nil
}

// calculateTestcasesSquareSum calculates square sum of the input testcases, and return a list of sum.
func calculateTestcasesSquareSum(reader *bufio.Reader, num int, result []int) ([]int, error) {
	if num == 0 {
		return result, nil
	}

	// 1. read length of testcase
	lengthTestcase, err := readSingleInteger(reader)
	if err != nil {
		return nil, err
	}

	// 2. read testcases
	testcase, err := readTestcase(reader, lengthTestcase)
	if err != nil {
		return nil, err
	}

	// 3, calculate square sum
	sum := calculateSquare(testcase, 0)

	result = append(result, sum)
	num = num - 1

	return calculateTestcasesSquareSum(reader, num, result)
}

// readTestcase reads integer array input of single testcase from standard input.
func readTestcase(reader *bufio.Reader, num int) ([]int, error) {
	if num == 0 {
		return nil, errors.New("invalid testcase length")
	}

	strTestCases, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	arrTestCases := strings.Fields(strTestCases)

	intTestCases, err := convertStrArrToIntArr(arrTestCases, []int{})
	if err != nil {
		return nil, err
	}

	if len(intTestCases) != num {
		return nil, errors.New("unmatch length of input testcase and given length")
	}

	return intTestCases, nil
}

// convertStrArrToIntArr converts input string array to an integer array
// use recursive solution instead of iterating with for loop.
func convertStrArrToIntArr(input []string, result []int) ([]int, error) {
	if len(input) == 0 {
		return result, nil
	}

	num, err := strconv.ParseInt(strings.TrimSpace(input[0]), 10, 32)
	if err != nil {
		return nil, err
	}

	result = append(result, int(num))
	input = input[1:]

	return convertStrArrToIntArr(input, result)
}

// calculateSquare converts calculates square sum of the input array
// use recursive solution instead of iterating with for loop.
func calculateSquare(input []int, result int) int {
	if len(input) == 0 {
		return result
	}

	if input[0] > 0 {
		result += input[0] * input[0]
	}

	input = input[1:]

	return calculateSquare(input, result)
}
