package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day1() int {
	file, err := os.Open("./day1/puzzle-input.txt")

	if err != nil {
		log.Fatal("open error := ", err)
	}
	defer file.Close()

	ans := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ans += solve(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("scanner error")
	}

	return ans
}

func solve(s string) int {
	left := 0
	right := 0

	// left side
	for i := 0; i < len(s); i++ {
		if int(s[i]-'0') > 0 && int(s[i]-'0') < 10 {
			left = int(s[i] - '0')
			break
		}
	}

	// right side
	for i := len(s) - 1; i >= 0; i-- {
		if int(s[i]-'0') > 0 && int(s[i]-'0') < 10 {
			right = int(s[i] - '0')
			break
		}
	}
	l := strconv.Itoa(left)
	r := strconv.Itoa(right)
	n, err := strconv.Atoi(l + r)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return n
}

// unicode.IsNumber(rune(s[i]))
