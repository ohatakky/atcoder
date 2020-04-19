package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	mod = 1e9 + 7
)

var dp = [100005][13]int{}

func main() {
	s := readString()
	n := len(s)

	dp[0][0] = 1
	for i := 0; i < n; i++ {
		var c int
		if string(s[i]) == "?" {
			c = -1
		} else {
			c = int(byte(s[i]) - '0')
		}

		for ii := 0; ii < 10; ii++ {
			if c != -1 && c != ii {
				continue
			}
			for iii := 0; iii < 13; iii++ {
				dp[i+1][(iii*10+ii)%13] = dp[i+1][(iii*10+ii)%13] + dp[i][iii]
			}
		}
		for ii := 0; ii < 13; ii++ {
			dp[i+1][ii] %= mod
		}
	}
	res := dp[n][5]
	fmt.Println(res)
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

/*-------------------init-------------------*/

var (
	readString func() string
)

func init() {
	readString = newReadString(os.Stdin)
}
