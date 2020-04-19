package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

const (
	INF_INT = 1000000000000000000
)

var (
	readString func() string
)

func init() {
	readString = newReadString(os.Stdin)
}

func main() {
	n := readf()
	N := int(n)

	maxMulti := math.Ceil(math.Sqrt(n))
	min := INF_INT
	a := 1
	for float64(a) <= maxMulti {
		if N%a == 0 {
			tmp := a + int(math.Ceil(float64(N)/float64(a))) - 2
			if min > tmp {
				min = tmp
			}
		}
		a = a + 1
	}

	fmt.Println(min)
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

func subL(list []int) []int {
	sl := make([]int, len(list)-1)
	for i := range list {
		if i == len(list)-1 {
			continue
		}
		sub := list[i+1] - list[i]
		sl[i] = sub
	}
	return sl
}
