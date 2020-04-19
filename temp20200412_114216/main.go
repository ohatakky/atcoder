package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type IsLand map[int][]int

var (
	n        int
	lands    = make(IsLand)
	visited  = make(map[int]bool)
	maxDepth = 2
	POSSIBLE bool
)

func main() {
	n = readInt()
	m := readInt()

	for i := 0; i < m; i++ {
		a := readInt()
		b := readInt()
		lands[a] = append(lands[a], b)
	}

	dfs(1, 1)
	if POSSIBLE {
		fmt.Println("POSSIBLE")
	} else {
		fmt.Println("IMPOSSIBLE")
	}
}

func dfs(p, depth int) {
	if depth >= 3 {
		return
	}
	visited[p] = true
	dests, ok := lands[p]
	if !ok {
		return
	}
	for _, d := range dests {
		if visited[d] {
			continue
		}
		if d == n {
			POSSIBLE = true
		}
		dfs(d, depth+1)
	}
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

/*-------------------init-------------------*/

var (
	readString func() string
)

func init() {
	readString = newReadString(os.Stdin)
}
