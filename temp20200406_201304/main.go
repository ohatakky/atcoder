package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readf(sc *bufio.Scanner) float64 {
	f, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func Abs(x int) int {
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

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	x := make([]int, m)
	for i := range x {
		sc.Scan()
		x[i], _ = strconv.Atoi(sc.Text())
	}

	if m == 1 || m < n {
		fmt.Println(0)
		return
	}

	sort.Ints(x)
	sl := subL(x)

	c := Abs(x[len(x)-1] - x[0])
	for i := 0; i < n-1; i++ {
		c = c - sl[len(sl)-1-i]
	}

	fmt.Println(c)
}
