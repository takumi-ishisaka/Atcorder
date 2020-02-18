package main

// -----問題文-----
// N個のサイコロが左から右に一列に並べてあります。左からi番目のサイコロは
// 1からpiまでのpi種類の目がそれぞれ等確率で出ます。
// 隣接するK個のサイコロを選んでそれぞれ独立に振ったとき、出る目の合計の期待値の最大値を求めてください。
// -----制約-----
// 1≤K≤N≤200000
// 1≤pi≤1000
// 入力で与えられる値は全て整数
// -----入力-----
// 入力は以下の形式で標準入力から与えられる。
// N K
// p1 ... pN
// -----出力-----
// 隣接するK個のサイコロを選んで振ったときに出る目の合計の期待値の最大値を出力せよ。
// なお、想定解答との絶対誤差または相対誤差が10^−6以下であれば正解として扱われる。

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var stdin = func() *bufio.Scanner {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Split(bufio.ScanWords)
	return stdin
}()

func main() {
	// 入力
	N := intConvert(readString())
	K := intConvert(readString())

	// 計算量O(N) 期待値スライスの作成
	var piExpectValues []float64
	for i := 0; i < N; i++ {
		pi := intConvert(readString())
		piExpectValue := expectValueCalculation(pi)
		piExpectValues = append(piExpectValues, piExpectValue)
	}

	// 期待値の論理和を計算する。
	var piExpectValueCumulativeSums []float64
	piExpectValueCumulativeSums = append(piExpectValueCumulativeSums, piExpectValues[0])
	for i := 1; i < len(piExpectValues); i++ {
		cumulativSum := piExpectValueCumulativeSums[i-1] + piExpectValues[i]
		piExpectValueCumulativeSums = append(piExpectValueCumulativeSums, cumulativSum)
	}

	// 隣接するK個の期待値の合計最大値を計算する
	expectValueMaxSum := piExpectValueCumulativeSums[0]
	if N == K {
		expectValueMaxSum = piExpectValueCumulativeSums[N-1]
	} else {

		for i := 1; i < N-K; i++ {
			priExpectValueMaxSum := piExpectValueCumulativeSums[i+K] - piExpectValueCumulativeSums[i]
			if expectValueMaxSum < priExpectValueMaxSum {
				expectValueMaxSum = priExpectValueMaxSum
			}
		}
	}

	// 出力
	fmt.Printf("%f", expectValueMaxSum)
}

// readString() スペース区切りで読み込んだ文字列を返す
func readString() string {
	stdin.Scan()
	return stdin.Text()
}

// intConver(string) 文字列をintに直す
func intConvert(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

// expectValueCalculation() 1からPまでのP種類の目が等確率で出るサイコロの出目の期待値を計算する
func expectValueCalculation(num int) float64 {
	var priExpectValue float64
	priExpectValue = (1.0 + float64(num)) / 2.0
	return priExpectValue
}
