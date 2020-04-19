package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readf(sc *bufio.Scanner) float64 {
	f, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	a := make([]int, 2)
	sc.Scan()
	a[0], _ = strconv.Atoi(sc.Text())
	sc.Scan()
	a[1], _ = strconv.Atoi(sc.Text())
	b := make([]int, 2)
	sc.Scan()
	b[0], _ = strconv.Atoi(sc.Text())
	sc.Scan()
	b[1], _ = strconv.Atoi(sc.Text())
	c := make([]int, 2)
	sc.Scan()
	c[0], _ = strconv.Atoi(sc.Text())
	sc.Scan()
	c[1], _ = strconv.Atoi(sc.Text())

	a[0] = a[0] - c[0]
	a[1] = a[1] - c[1]
	b[0] = b[0] - c[0]
	b[1] = b[1] - c[1]
	c[0] = c[0] - c[0]
	c[1] = c[1] ^ c[1]

	res := float64(AbsInt((a[0]*b[1] - a[1]*b[0]))) / 2.0
	fmt.Println(res)
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
