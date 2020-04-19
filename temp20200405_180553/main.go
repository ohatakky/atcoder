package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Record struct {
	ID     int
	Number int
}

type Records []Record

func (r Records) Len() int {
	return len(r)
}

func (r Records) Less(i, j int) bool {
	return r[i].Number < r[j].Number
}

func (r Records) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

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

	a := make(Records, n)
	for i := range a {
		sc.Scan()
		a[i].ID = i + 1
		a[i].Number, _ = strconv.Atoi(sc.Text())
	}

	sort.Sort(Records(a))

	for i := range a {
		fmt.Printf("%d ", a[i].ID)
	}
	fmt.Println()
}
