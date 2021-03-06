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

func main() {
	n := readi()
	k := readi()
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = readi()
	}

	mapSum := make(map[int]int)
	mapSum[0] = p[0]
	for i := 1; i < n; i++ {
		mapSum[i] = p[i] + mapSum[i-1]
	}

	maxIdx := 0
	maxSum := 0
	for i := k - 1; i < n; i++ {
		tmp := mapSum[i] - mapSum[i-k]
		if maxSum < tmp {
			maxSum = tmp
			maxIdx = i
		}
	}

	ans := 0.0
	for i := maxIdx - k + 1; i <= maxIdx; i++ {
		ans += expectedValue(p[i])
	}

	fmt.Printf("%f\n", ans)
}

func expectedValue(p int) float64 {
	res := 0.0
	for i := 1; i <= p; i++ {
		res += float64(i) * (1 / float64(p))
	}
	return res
}

func max64(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
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

// func chmin(x *int, y int) bool {
// 	if *x > y {
// 		*x = y
// 		return true
// 	}
// 	return false
// }

// func chmax(x *int, y int) bool {
// 	if *x < y {
// 		*x = y
// 		return true
// 	}
// 	return false
// }

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func permutations(arr []int) [][]int {
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
