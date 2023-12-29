package parser

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var numWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func ParseTxt(fn string) (int, error) {
	fp, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(fp)

	var nums []int

	for scanner.Scan() {
		str := scanner.Text()
		first := searchFwd(str)
		last := searchBwd(str)

		found, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, found)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fp.Close()

	return sum(nums), nil
}

func searchFwd(str string) string {
	for i := 0; i < len(str); i++ {
		if _, err := strconv.Atoi(string(str[i])); err == nil {
			return string(str[i])
		}
	}
	return "0"
}

func searchBwd(str string) string {
	for i := len(str) - 1; i >= 0; i-- {
		if _, err := strconv.Atoi(string(str[i])); err == nil {
			return string(str[i])
		}
	}
	return "0"
}

func sum(numArr []int) int {
	sum := 0
	for _, val := range numArr {
		sum += val
	}

	return sum
}
