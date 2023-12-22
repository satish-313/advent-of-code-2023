package day2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var m = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Day2_1() int {
	f, err := os.Open("./day2/input1.txt")
	if err != nil {
		log.Fatal("os Open error ", err)
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	sum := 0
	sum1 := 0
	for scan.Scan() {
		sum1 += parse(scan.Text())
		sum += solveLater(scan.Text())
	}
	if err := scan.Err(); err != nil {
		log.Fatal("scan error ", err)
	}
	return sum
}

func solveLater(s string) int {
	p := strings.Split(s, ":")
	later := strings.Split(p[1], ";")

	var red, green, blue int

	for i := 0; i < len(later); i++ {
		items := strings.Split(later[i], ",")
		for j := 0; j < len(items); j++ {
			item := strings.Split(items[j], " ")

			colour := item[2]
			value, err := strconv.Atoi(item[1])
			if err != nil {
				return 0
			}

			if colour == "red" && value > red {
				red = value
			} else if colour == "blue" && value > blue {
				blue = value
			} else if colour == "green" && value > green {
				green = value
			}
		}
	}

	return red * blue * green
}

func parse(s string) int {
	p := strings.Split(s, ":")
	id, err := strconv.Atoi(strings.Split(p[0], " ")[1])
	if err != nil {
		return 0
	}
	if !parseLater(p[1]) {
		return 0
	}
	return id
}

func parseLater(s string) bool {
	later := strings.Split(s, ";")

	for i := 0; i < len(later); i++ {
		items := strings.Split(later[i], ",")
		for j := 0; j < len(items); j++ {
			item := strings.Split(items[j], " ")

			colour := item[2]
			value, err := strconv.Atoi(item[1])
			if err != nil {
				return false
			}

			val := m[colour]

			if value > val {
				return false
			}
		}
	}
	return true
}
