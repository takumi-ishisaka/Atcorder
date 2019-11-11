package main

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
	var output int
	// a<= √Nとなるため
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

	fmt.Println(output)
}
