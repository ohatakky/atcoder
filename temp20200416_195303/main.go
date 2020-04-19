package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var (
	readString func() string
)

func init() {
	readString = newReadString(os.Stdin)
}

type sortString []string

func (s sortString) Len() int {
	return len(s)
}

func (s sortString) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s sortString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func main() {
	n := readInt()
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = readString()
	}

	// 1
	sort.Sort(sortString(s))

	// 2
	mapArray := make([]map[string]int, n)
	for i := range mapArray {
		mapArray[i] = make(map[string]int)
	}

	for i := range s {
		for ii := range s[i] {
			mapArray[i][string(s[i][ii])] = mapArray[i][string(s[i][ii])] + 1
		}
	}

	// 3
	tmp := make(map[string]int)
	for i := range mapArray[0] {
		tmp[i] = mapArray[0][i]
		for ii := 1; ii < n; ii++ {
			if tmp[i] > mapArray[ii][i] {
				tmp[i] = mapArray[ii][i]
			}
		}
	}

	ans := ""
	for i := range tmp {
		for ii := 0; ii < tmp[i]; ii++ {
			ans = ans + i
		}
	}

	fmt.Println(SortString(ans))
}

/*-------------------inputs-------------------*/

func readInt() int {
	return int(readInt64())
}

func readInt64() int64 {
	i, err := strconv.ParseInt(readString(), 0, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
}

func readf() float64 {
	f, err := strconv.ParseFloat(readString(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func newReadString(ior io.Reader) func() string {
	r := bufio.NewScanner(ior)
	r.Buffer(make([]byte, 1024), int(1e+11))
	r.Split(bufio.ScanWords)

	return func() string {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Text()
	}
}

/*-------------------utilities-------------------*/

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
