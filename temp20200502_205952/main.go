package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

type St struct {
	c, d int
}

func main() {
	N := readi()
	M := readi()
	Q := readi()
	a := make([]int, Q)
	b := make([]int, Q)
	c := make([]int, Q)
	d := make([]int, Q)
	for i := 0; i < Q; i++ {
		a[i] = readi() - 1
		b[i] = readi() - 1
		c[i] = readi()
		d[i] = readi()
	}

	A := make([]int, N)
	A[0] = 1

	var ans int
	var dfs func(i int)
	dfs = func(i int) {
		if i == N {
			var t int
			for j := range a {
				if A[b[j]]-A[a[j]] == c[j] {
					t += d[j]
				}
			}
			ans = max(ans, t)
			return
		}
		for x := A[i-1]; x <= M; x++ {
			A[i] = x
			dfs(i + 1)
			A[i] = 0
		}
	}

	dfs(1)

	println(ans)
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

func factorical(n int) int {
	if n > 0 {
		return n * factorical(n-1)
	}
	return 1
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
