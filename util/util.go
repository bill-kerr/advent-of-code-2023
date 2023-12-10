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

func IntPow(base, exponent int) int {
	if exponent == 0 {
		return 1
	}

	result := base
	for i := 2; i <= exponent; i++ {
		result *= base
	}

	return result
}

func ParseInts(stringIntegers []string) []int {
	parsed := []int{}

	for _, str := range stringIntegers {
		parsedInt, _ := strconv.ParseInt(str, 10, 64)
		parsed = append(parsed, int(parsedInt))
	}

	return parsed
}

func Map[T, V any](ts []T, fn func(T, int) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t, i)
	}
	return result
}

func Filter[T any](ts []T, fn func(T, int) bool) []T {
	return Reduce(ts, func(cur T, acc []T, idx int) []T {
		if fn(cur, idx) {
			acc = append(acc, cur)
		}
		return acc
	}, []T{})
}

func Reduce[T, V any](ts []T, fn func(T, V, int) V, initial V) V {
	result := initial
	for i, t := range ts {
		result = fn(t, result, i)
	}
	return result
}
