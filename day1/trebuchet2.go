package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day1_2() int {
	file, err := os.Open("./day1/puzzle-input2.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ans := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ans += solve1(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("scanner error")
	}

	return ans
}

func solve1(s string) int {
	var left, right int

	// left
leftside:
	for i := 0; i < len(s); i++ {
		switch {
		case s[i]-'0' > 0 && s[i]-'0' < 10:
			left = int(s[i] - '0')
			break leftside
		case s[i] == 'o':
			if checkStr(i, &s, "one") {
				left = 1
				break leftside
			}
		case s[i] == 't':
			if checkStr(i, &s, "two") {
				left = 2
				break leftside
			}
			if checkStr(i, &s, "three") {
				left = 3
				break leftside
			}
		case s[i] == 'f':
			if checkStr(i, &s, "four") {
				left = 4
				break leftside
			}
			if checkStr(i, &s, "five") {
				left = 5
				break leftside
			}
		case s[i] == 's':
			if checkStr(i, &s, "seven") {
				left = 7
				break leftside
			}
			if checkStr(i, &s, "six") {
				left = 6
				break leftside
			}
		case s[i] == 'e':
			if checkStr(i, &s, "eight") {
				left = 8
				break leftside
			}
		case s[i] == 'n':
			if checkStr(i, &s, "nine") {
				left = 9
				break leftside
			}
		}

	}
rightside:
	for i := len(s) - 1; i >= 0; i-- {
		switch {
		case s[i]-'0' > 0 && s[i]-'0' < 10:
			right = int(s[i] - '0')
			break rightside
		case s[i] == 'o':
			if checkStr(i, &s, "one") {
				right = 1
				break rightside
			}
		case s[i] == 't':
			if checkStr(i, &s, "two") {
				right = 2
				break rightside
			}
			if checkStr(i, &s, "three") {
				right = 3
				break rightside
			}
		case s[i] == 'f':
			if checkStr(i, &s, "four") {
				right = 4
				break rightside
			}
			if checkStr(i, &s, "five") {
				right = 5
				break rightside
			}
		case s[i] == 's':
			if checkStr(i, &s, "seven") {
				right = 7
				break rightside
			}
			if checkStr(i, &s, "six") {
				right = 6
				break rightside
			}
		case s[i] == 'e':
			if checkStr(i, &s, "eight") {
				right = 8
				break rightside
			}
		case s[i] == 'n':
			if checkStr(i, &s, "nine") {
				right = 9
				break rightside
			}
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

func checkStr(i int, s *string, c string) bool {
	if len(*s)-i < len(c) {
		return false
	}

	for j := 0; j < len(c); j++ {
		if (*s)[i] != c[j] {
			return false
		}
		i++
	}

	return true
}
