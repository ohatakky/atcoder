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

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	d := make([]int, n)
	for i := range d {
		sc.Scan()
		d[i], _ = strconv.Atoi(sc.Text())
	}

	sort.Ints(d)

	half := len(d) / 2
	fmt.Println(d[half] - d[half-1])
}
