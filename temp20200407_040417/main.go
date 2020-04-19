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

func permutation(n int, k int) int {
	v := 1
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			v *= (n - i)
		}
	} else if k > n {
		v = 0
	}
	return v
}

func factorial(n int) int {
	return permutation(n, n-1)
}

func combination(n int, k int) int {
	return permutation(n, k) / factorial(k)
}

/*
異なる整数が5個与えられます。
この中から3つ選んでその和で表すことの出来る整数のうち、3番目に大きいものを出力してください。
*/
// Note: 10通り

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	num := 5
	x := make([]int, num)
	for i := 0; i < num; i++ {
		sc.Scan()
		x[i], _ = strconv.Atoi(sc.Text())
	}

	sort.Ints(x)

	a1 := x[4] + x[3] + x[0]
	a2 := x[4] + x[2] + x[1]
	if a1 > a2 {
		fmt.Println(a1)
		return
	}
	fmt.Println(a2)
}
