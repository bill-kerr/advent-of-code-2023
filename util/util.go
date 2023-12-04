package util

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
)

func OpenAndRead(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to read text file")
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	lines = []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func SumSlice(slice []int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}
	return sum
}

func Reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Atoi(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("Failed to convert string to integer")
	}
	return val
}

func Rtoi(r rune) int {
	return int(r - '0')
}

func RtoDigit(r rune) (int, error) {
	digit := Rtoi(r)

	if digit < 0 || digit > 9 {
		return digit, errors.New("provided rune not a valid digit")
	}

	return digit, nil
}
