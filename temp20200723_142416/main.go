package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type vacation struct {
	a int
	b int
	c int
}

func main() {
	n := readi()
	v := make([]vacation, n)
	for i := 0; i < n; i++ {
		v[i].a = readi()
		v[i].b = readi()
		v[i].c = readi()
	}

	// 0: a, 1: b, 2: c
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 3)
	}

	dp[0][0] = v[0].a
	dp[0][1] = v[0].b
	dp[0][2] = v[0].c
	for i := 1; i < n; i++ {
		for j := 0; j < 3; j++ {
			switch j {
			case 0:
				tmp1 := max(dp[i-1][1], dp[i-1][1]+v[i].a)
				tmp2 := max(dp[i-1][2], dp[i-1][2]+v[i].a)
				dp[i][0] = max(tmp1, tmp2)
			case 1:
				tmp1 := max(dp[i-1][0], dp[i-1][0]+v[i].b)
				tmp2 := max(dp[i-1][2], dp[i-1][2]+v[i].b)
				dp[i][1] = max(tmp1, tmp2)
			case 2:
				tmp1 := max(dp[i-1][0], dp[i-1][0]+v[i].c)
				tmp2 := max(dp[i-1][1], dp[i-1][1]+v[i].c)
				dp[i][2] = max(tmp1, tmp2)
			}
		}
	}

	ans := 0
	for i := 0; i < 3; i++ {
		ans = max(ans, dp[n-1][i])
	}

	fmt.Println(ans)
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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func nextPermutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeFactors(n int) []int {
	factors := make([]int, 0)
	i := 2
	for i*i <= n {
		r := n % i
		if r != 0 {
			i++
		} else if r == 0 {
			n /= i
			factors = append(factors, i)
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

func divisor(n int) (res []int) {
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			res = append(res, i)
			if i*i != n {
				res = append(res, n/i)
			}
		}
	}
	sort.Ints(res)
	return
}

func searchInts(a []int, k int) int {
	return binarySearch(len(a), func(i int) bool {
		return a[i] >= k
	})
}

func binarySearch(n int, f func(int) bool) int {
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if !f(h) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func reverseStr(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteByte(s[i])
	}
	return b.String()
}

type Queue struct {
	items []int
}

func (q *Queue) push(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) pop() {
	q.items = q.items[1:]
}

func (q *Queue) front() int {
	return q.items[0]
}

func (q *Queue) size() int {
	return len(q.items)
}

type Graph struct {
	n     int
	edges [][]int
}

func newGraph(n int) *Graph {
	g := &Graph{
		n:     n,
		edges: make([][]int, n),
	}
	return g
}

func (g *Graph) addEdge(u, v int) {
	g.edges[v] = append(g.edges[v], u)
	g.edges[u] = append(g.edges[u], v)
}

func dfs(c int, edges [][]int, visited map[int]struct{}) {
	visited[c] = struct{}{}

	for _, v := range edges[c] {
		_, flag := visited[v]
		if flag {
			continue
		}
		dfs(v, edges, visited)
	}
}

func bfs(c int, graph *Graph) {
	next := make([]int, 0)
	next = append(next, c)
	visited := make(map[int]struct{})

	for len(next) != 0 {
		u := next[0]
		next = next[1:]
		visited[u] = struct{}{}

		for _, v := range graph.edges[u] {
			_, flag := visited[v]
			if flag {
				continue
			}

			// なにか処理

			next = append(next, v)
		}
	}
}

/*-------------------init-------------------*/

const (
	INF = 1000000000000000000
	MOD = 1e9 + 7
)

var (
	readString func() string
)

func init() {
	readString = newReadString(os.Stdin)
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

func readiSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = readi()
	}
	return slice
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
