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

// TODO: 市の順番に出力できるデータ構造にする
type Prefectures map[int]Years

type Cities []City

type City struct {
	Name int
	ID   int
}

type Years []int

func (c Years) Len() int {
	return len(c)
}

func (c Years) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c Years) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	prefectures := make(Prefectures, n)
	for i := 0; i < m; i++ {
		sc.Scan()
		p, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		y, _ := strconv.Atoi(sc.Text())
		prefectures[p] = append(prefectures[p], y)
	}

	for k, years := range prefectures {
		sort.Sort(Years(years))
		for i := range years {
			fmt.Printf("%06d%06d\n", k, i+1)
		}
	}
}
