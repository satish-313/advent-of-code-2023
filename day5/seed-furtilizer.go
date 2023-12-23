package day5

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	seed     []uint64
	soil     [][3]uint64
	fer      [][3]uint64
	water    [][3]uint64
	light    [][3]uint64
	temp     [][3]uint64
	humd     [][3]uint64
	location [][3]uint64
)

func Day5() uint64 {
	file, err := os.Open("./day5/input.txt")
	if err != nil {
		log.Fatal("file error", err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	count := 0
	for s.Scan() {
		str := s.Text()
		if len(str) == 0 {
			count++
			continue
		}
		if count == 0 {
			seed = parseSeed(str)
		} else if count == 1 {
			if !strings.HasSuffix(str, ":") {
				soil = append(soil, [3]uint64(parseNumber(str)))
			}
		} else if count == 2 {
			if !strings.HasSuffix(str, ":") {
				fer = append(fer, [3]uint64(parseNumber(str)))
			}
		} else if count == 3 {
			if !strings.HasSuffix(str, ":") {
				water = append(water, [3]uint64(parseNumber(str)))
			}
		} else if count == 4 {
			if !strings.HasSuffix(str, ":") {
				light = append(light, [3]uint64(parseNumber(str)))
			}
		} else if count == 5 {
			if !strings.HasSuffix(str, ":") {
				temp = append(temp, [3]uint64(parseNumber(str)))
			}
		} else if count == 6 {
			if !strings.HasSuffix(str, ":") {
				humd = append(humd, [3]uint64(parseNumber(str)))
			}
		} else if count == 7 {
			if !strings.HasSuffix(str, ":") {
				location = append(location, [3]uint64(parseNumber(str)))
			}
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal("s error", err)
	}

	lopper := []*[][3]uint64{&soil, &fer, &water, &light, &temp, &humd, &location}

	seedRange := make([]uint64, 0)
	for i := 0; i < len(seed); {
		start := seed[i]
		ran := seed[i+1]
		var j uint64
		for j = 0; j < ran; j++ {
			seedRange = append(seedRange, start+j)
		}
		i += 2
	}
	// fmt.Pruint64ln(seedRange)
	// for idx, val := range seed {
	// 	solve(idx, val, &seed, &soil)
	// }
	// for idx, val := range seed {
	// 	solve(idx, val, &seed, &fer)
	// }
	// for idx, val := range seed {
	// 	solve(idx, val, &seed, &water)
	// }
	// for idx, val := range seed {
	// 	solve(idx, val, &seed, &light)
	// }
	// for idx, val := range seed {
	// 	solve(idx, val, &seed, &temp)
	// }
	// for idx, val := range seed {
	// 	solve(idx, val, &seed, &humd)
	// }
	// for idx, val := range seed {
	// 	solve(idx, val, &seed, &location)
	// }
	for _, loop := range lopper {
		for idx, val := range seedRange {
			solve(uint64(idx), val, &seedRange, loop)
		}
	}
	// fmt.Pruint64ln(seedRange)
	return findMin(&seedRange)
}

func findMin(arr *[]uint64) uint64 {
	min := uint64(math.MaxUint64)
	for _, val := range *arr {
		if min > val {
			min = val
		}
	}
	return min
}

func solve(idx, value uint64, ans *[]uint64, next *[][3]uint64) {
	for _, val := range *next {
		d := val[0]
		s := val[1]
		r := val[2]

		if value >= s && value <= s+r {
			diff := uint64(math.Abs(float64(value - s)))
			(*ans)[idx] = d + diff
			return
		}
	}
	(*ans)[idx] = value
}
func parseSeed(s string) []uint64 {
	p1 := strings.Split(s, ":")
	seeds := strings.Split(p1[1], " ")
	seed := make([]uint64, 0)
	for _, v := range seeds {
		if len(strings.TrimSpace(v)) > 0 {
			num, _ := strconv.ParseUint(v, 10, 64)
			seed = append(seed, num)
		}
	}
	return seed
}

func parseNumber(s string) [3]uint64 {
	seeds := strings.Split(s, " ")
	var seed [3]uint64
	idx := 0
	for _, v := range seeds {
		if len(strings.TrimSpace(v)) > 0 {
			num, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			seed[idx] = num
			idx++
		}
	}
	return seed
}
