package day3

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day3() int {
	file, err := os.Open("./day3/input1.txt")
	if err != nil {
		log.Fatal("file error ", err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	matrix := make([][]string, 0)

	for scan.Scan() {
		line := scan.Text()
		row := strings.Split(line, "")
		matrix = append(matrix, row)
	}
	if err := scan.Err(); err != nil {
		log.Fatal("scan error ", err)
	}
	sum := solve1(&matrix)
	return sum
}

func solve1(m *[][]string) int {
	sum := 0
	row := len((*m))
	col := len((*m)[0])

	dir := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			char := makeRune((*m)[r][c])
			if char != '*' {
				continue
			}
			count := 0
			tsum := 1
			for d := 0; d < len(dir); d++ {
				i := r + dir[d][0]
				j := c + dir[d][1]

				// outof matrix
				if i < 0 || i > row || j < 0 || j > col {
					continue
				}

				if isNumber(makeRune((*m)[i][j])) {
					tsum *= findNumber(&(*m)[i], j)
					count++
				}
			}
			if count == 2 {
				sum += tsum
			}
		}
	}
	return sum
}

func findNumber(row *[]string, j int) int {
	var builder strings.Builder
	var i int
	for i = j; i >= 0 && i < len((*row)); i-- {
		if isNumber(makeRune((*row)[i])) {
			continue
		} else {
			break
		}
	}

	j = i + 1
	for j < len(*row) {
		if isNumber(makeRune((*row)[j])) {
			builder.WriteString((*row)[j])
			(*row)[j] = string('.')
			j++
		} else {
			break
		}
	}

	num, err := strconv.Atoi(builder.String())
	if err != nil {
		log.Fatal("convert error", err)
	}
	return num
}

func isNumber(r rune) bool {
	if r-'0' >= 0 && r-'0' < 10 {
		return true
	}
	return false
}

func makeRune(s string) rune {
	var r rune
	for _, val := range s {
		r = val
	}
	return r
}
