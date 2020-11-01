package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type sousa struct {
	kind     int
	op1, op2 int
}

type status int

const (
	N status = iota
	Y
	Ykari
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func solve() {
	n, q := getInt(), getInt()

	// Q行分のフォローログを構造体に記憶
	ss := make([]sousa, q)
	for i := range ss {
		ss[i].kind = getInt()
		ss[i].op1 = getInt() - 1
		if ss[i].kind == 1 {
			ss[i].op2 = getInt() - 1
		}
	}

	// 出力用の解答を保存するスライスの初期化
	follows := make([][]status, n)
	for i := range follows {
		follows[i] = make([]status, n)
		for j := range follows[i] {
			follows[i][j] = N
		}
	}

	for _, s := range ss {
		switch s.kind {
		case 1:
			// フォロー
			follows[s.op1][s.op2] = Y
		case 2:
			// 全フォロー返し
			for i := 0; i < n; i++ {
				if follows[i][s.op1] == Y {
					follows[s.op1][i] = Y
				}
			}
		case 3:
			// フォローフォロー
			for i := 0; i < n; i++ {
				if follows[s.op1][i] != Y {
					continue
				}
				// iはop1がフォローしているユーザー
				for j := 0; j < n; j++ {
					// iがjをフォローしていて、jはop1ではなく、op1はjをフォローしていない場合、op1はjをフォローする
					if follows[i][j] == Y && follows[s.op1][j] == N && j != s.op1 {
						follows[s.op1][j] = Ykari
					}
				}
			}
			// フォローフォローが必要な相手がわかったので、仮YフラグをYにする
			for i := range follows[s.op1] {
				if follows[s.op1][i] == Ykari {
					follows[s.op1][i] = Y
				}
			}
		}
	}

	for i := range follows {
		for _, f := range follows[i] {
			if f == Y {
				fmt.Fprint(wr, "Y")
			} else {
				fmt.Fprint(wr, "N")
			}
		}
		fmt.Fprintln(wr, "")
	}
}

func main() {
	maxBufSize := int(1e8)
	sc.Buffer(make([]byte, 4096), maxBufSize)
	sc.Split(bufio.ScanWords)
	solve()
	wr.Flush()
}

func getInt() (ret int) {
	sc.Scan()
	ret, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return
}
