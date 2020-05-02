package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type start struct {
	x int
	y int
}

var (
	h       int
	w       int
	c       [][]string
	visited [][]bool
)

func main() {
	h = readi()
	w = readi()
	c = make([][]string, h)
	s := new(start)
	for i := 0; i < h; i++ {
		c[i] = make([]string, w)
		tmp := readString()
		c[i] = strings.Split(tmp, "")
		for ii := 0; ii < w; ii++ {
			if c[i][ii] == "s" {
				s.x = i
				s.y = ii
			}
		}
	}

	visited = make([][]bool, h)
	for i := 0; i < h; i++ {
		visited[i] = make([]bool, w)
	}

	if dfs(s.x, s.y) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func dfs(i, ii int) bool {
	if i <= -1 || i >= h || ii <= -1 || ii >= w {
		return false
	}
	if c[i][ii] == "#" {
		return false
	}
	if visited[i][ii] {
		return false
	}
	if c[i][ii] == "g" {
		return true
	}

	visited[i][ii] = true
	if dfs(i, ii+1) {
		return true
	}
	if dfs(i+1, ii) {
		return true
	}
	if dfs(i, ii-1) {
		return true
	}
	if dfs(i-1, ii) {
		return true
	}

	return false
}

/*-------------------utilities-------------------*/

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func chmin(x *int, y int) bool {
	if *x > y {
		*x = y
		return true
	}
	return false
}

func chmax(x *int, y int) bool {
	if *x < y {
		*x = y
		return true
	}
	return false
}

// 最大公約数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// 最小公倍数
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func comb(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	return fac[n] * (finv[k] * finv[n-k] % MOD) % MOD
}

/*-------------------init-------------------*/

const (
	INF = 1000000000000000000
	MOD = 1e9 + 7

	COMB_MAX = 510000
)

var (
	readString func() string
)

var (
	fac  = [COMB_MAX]int{}
	finv = [COMB_MAX]int{}
	inv  = [COMB_MAX]int{}
)

func init() {
	readString = newReadString(os.Stdin)
}

func combInit() {
	fac[0] = 1
	fac[1] = 1
	finv[0] = 1
	finv[1] = 1
	inv[1] = 1
	for i := 2; i < COMB_MAX; i++ {
		fac[i] = fac[i-1] * i % MOD
		inv[i] = MOD - inv[MOD%i]*(MOD/i)%MOD
		finv[i] = finv[i-1] * inv[i] % MOD
	}
}

/*-------------------inputs-------------------*/

func readi() int {
	return int(readi64())
}

func readi64() int64 {
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
