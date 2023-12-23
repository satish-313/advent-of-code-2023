package day5

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Range struct {
	source, target, len uint64
}

type RangeList []Range

var seeds []uint64 = []uint64{}
var maps [7]RangeList = [7]RangeList{}

func (r *RangeList) Insert(s, t, l uint64) {
	index := sort.Search(len(*r), func(i int) bool { return (*r)[i].source > s })
	*r = append(*r, Range{})
	copy((*r)[index+1:], (*r)[index:])
	(*r)[index] = Range{s, t, l}
}

func (r *RangeList) Find(s uint64) uint64 {
	index := sort.Search(len(*r), func(i int) bool { return (*r)[i].source > s })
	if index == 0 || s > (*r)[index-1].source+(*r)[index-1].len {
		return s
	}

	return (*r)[index-1].target + s - (*r)[index-1].source
}

func Day5a() int {
	file, err := os.Open("./day5/test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	l := scanner.Text()
	re := regexp.MustCompile(`\d+`) //number
	matches := re.FindAllString(l, -1)

	for _, m := range matches {
		v, _ := strconv.ParseUint(m, 10, 64)
		seeds = append(seeds, v)
	}
	scanner.Scan()

	for i := 0; i < 7; i++ {
		scanner.Scan()
		for scanner.Scan() {
			l = scanner.Text()
			if l == "" {
				break
			}

			matches = re.FindAllString(l, -1)
			v1, _ := strconv.ParseUint(matches[0], 10, 64)
			v2, _ := strconv.ParseUint(matches[1], 10, 64)
			v3, _ := strconv.ParseUint(matches[2], 10, 64)

			maps[i].Insert(v2, v1, v3)
		}
	}
	part1()
	part2()
	fmt.Println(maps)
	return 0
}

func part1() {
	s := uint64(math.MaxUint64)

	for _, seed := range seeds {
		v := seed
		for i := 0; i < 7; i++ {
			v = maps[i].Find(v)
		}
		if v < s {
			s = v
		}
	}
	fmt.Printf("%d\n", s)
}

func part2() {
	s := uint64(math.MaxUint64)

	allRanges := RangeList{}
	for i := 0; i < 7; i++ {
		for _, r := range maps[i] {
			allRanges.Insert(r.source, r.target, r.len)
		}
	}

	for i := 0; i < len(seeds); i += 2 {
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			v := seed
			for i := 0; i < 7; i++ {
				v = maps[i].Find(v)
			}
			if v < s {
				s = v
			}

			index := sort.Search(len(allRanges), func(i int) bool { return allRanges[i].source > seed })
			if index != len(allRanges) {
				seed = allRanges[index].source - 1
			}
		}
	}
	fmt.Printf("%d\n", s)
}
