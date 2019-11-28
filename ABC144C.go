package main

// -----問題文-----
// 高橋君は無限に広い掛け算表の上にいます。
// 掛け算表のマス（i，j）には整数i×jが書かれており、高橋君は最初（1，1）にいます。
// 高橋君は1回の移動で（i，j）から（i+1，j）か（i，j+1）のどちらかにのみ移ることができます。
// 整数Nが与えられるので、Nがかかれているますに到達するまでに必要な移動回数の最小値を求めてください。
// -----制約-----
// 2≦N≦10^12
// Nは整数である。

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// 入力文字列の読み込み
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	input, err := strconv.Atoi(stdin.Text())
	if err != nil {
		panic(err)
	}
	// （1，1）から（a，b）に至るまでの移動回数はa+b-2。したがって、問題は
	//「a×b=Nとなるような（a，b）について、a+b-2の最小値を求めよ」となる。
	// 対称性からa≦bとしてよいので、a<= √Nまでを全探索することにより、O（N）で解ける。
	// 今回は逆に√Nを求めてから、大きいほうから小さいほうに探索を行い、探索時間を減らした。
	// さらに素数判定を行うことでさらに探索を減らした。
	var output int
	// 素数判定
	prime := PrimeCheck(input)
	if prime {
		output = input - 1
	} else {
		aLimit := int(math.Sqrt(float64(input)))
		a := 1
		b := 1
		for i := aLimit; i > 0; i-- {
			if input%i == 0 {
				a = i
				b = input / i
				break
			}
		}
		if b == 1 {
			a = input
		}
		output = a + b - 2
	}
	fmt.Println(output)
}

// 素数判定
func PrimeCheck(input int) bool {
	if input == 1 {
		return false
	}
	if input == 2 {
		return true
	}
	if input%2 == 0 {
		return false
	}

	primeFlag := true
	rootInput := int(math.Sqrt(float64(x)))
	i := 3
	for i <= rootInput {
		if input%i == 0 {
			primeFlag = false
			break
		}
		i += 2
	}
	return primeFlag
}
