package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/fatih/color"
)

var (
	n int
	m int
	a []int
	b []int

	graph   map[int][]int
	visited []int

	ans int
)

func main() {
	n = readi()
	m = readi()
	a = make([]int, m)
	b = make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = readi()
		b[i] = readi()
	}

	// init graph
	graph = make(map[int][]int)
	for i := 0; i < m; i++ {
		graph[a[i]] = append(graph[a[i]], b[i])
		graph[b[i]] = append(graph[b[i]], a[i])
	}

	// init visited
	visited = make([]int, n+1)
	for i := 1; i <= n; i++ {
		visited[i] = -1
	}

	dfs(1, visited)
	fmt.Println(ans)
}

func dfs(depth int, visited []int) {
	color.Yellow("%d", depth)
	if visited[depth] != -1 {
		// return false
	}
	if depth == n {
		// return true
		ans++
	}
	visited[depth] = 1

	// return dfs(depth+1, visited)
	dfs(depth+1, visited)
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
