package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type vacation struct {
	a int
	b int
	c int
}

func main() {
	n := readi()
	vacations := make([]vacation, n)
	for i := 0; i < n; i++ {
		vacations[i].a = readi()
		vacations[i].b = readi()
		vacations[i].c = readi()
	}

	dp := make([][3]int, n)
	dp[0][0] = vacations[0].a
	dp[0][1] = vacations[0].b
	dp[0][2] = vacations[0].c
	for i := 1; i < n; i++ {
		for ii := 0; ii < 3; ii++ {
			for iii := 0; iii < 3; iii++ {
				if iii != ii {
					switch ii {
					case 0:
						chmax(&dp[i][0], vacations[i].a+dp[i-1][iii])
					case 1:
						chmax(&dp[i][1], vacations[i].b+dp[i-1][iii])
					case 2:
						chmax(&dp[i][2], vacations[i].c+dp[i-1][iii])
					}
				}
			}
		}
	}

	fmt.Println(max(max(dp[n-1][0], dp[n-1][1]), dp[n-1][2]))
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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
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
