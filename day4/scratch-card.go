package day4

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Day4() int {
	file, err := os.Open("./day4/input1.txt")
	if err != nil {
		log.Fatal("file error", err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	len := 0

	for s.Scan() {
		len += 1
	}
	if err := s.Err(); err != nil {
		log.Fatal("s error", err)
	}
	file.Seek(0, 0)
	scan := bufio.NewScanner(file)

	hashMap := make([]int, len)
	count := 0
	for scan.Scan() {
		temp := solve(scan.Text())
		if temp != 0 {
			hashMap[count]++
			for i := count + 1; i < len && i <= temp+count; i++ {
				hashMap[i] += hashMap[count]
			}
		} else {
			hashMap[count]++
		}
		// fmt.Println(hashMap)
		count += 1
	}
	if err := scan.Err(); err != nil {
		log.Fatal("scanner error", err)
	}

	sum := 0
	for _, val := range hashMap {
		sum += val
	}
	return sum
}

func solve(s string) int {
	count := 0
	winnerNum, yourNum := parser(&s)
	m := make(map[string]bool)

	for _, val := range winnerNum {
		if val != string("") {
			m[val] = true
		}
	}
	for _, val := range yourNum {
		if val != string("") && m[val] {
			count++
		}
	}

	// if count == 0 {
	// 	return 0
	// }
	// return power(2, count-1)

	return count
}

// func power(base, exponent int) int {
// 	result := 1
// 	for i := 0; i < exponent; i++ {
// 		result *= base
// 	}
// 	return result
// }

func parser(s *string) ([]string, []string) {
	p1 := strings.Split((*s), ":")
	p2 := strings.Split(p1[1], "|")

	wn := strings.Split(p2[0], " ")
	yn := strings.Split(p2[1], " ")

	return wn, yn
}
