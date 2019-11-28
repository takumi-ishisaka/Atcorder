package main

// -----問題文-----
// N頂点の木Gが与えられます。頂点には1からNまでの番号がついており、
// i本目の辺は頂点aiと頂点biを結んでいます。
// Gの辺を何色かで塗り分けることを考えます。 このとき、各頂点について、
// その頂点を端点に持つ辺の色がすべて相異なるようにしたいです。
// 上記の条件を満たす塗り分けの中で、使用する色の数が最小であるようなものを
// 1つ構築してください。
// -----制約-----
// 2≦N≦10^5
// 1≦ai＜bi≦N
// 入力は全て整数
// 与えられるグラフは木である

// -----入力-----
// 入力は以下の形式で標準入力から与えられる。
// N
// a1 b1
// a2 b2
//   ⋮
// aN−1 bN−1

// -----出力-----
// 出力はN行からなる
// 1行目のに使用する色の数Kを出力せよ
// i+1（1≦i≦N-1）にi番目の辺の色を表す整数ciを出力せよ
// ここで1≦ci≦kでなければならない
// 問題分中の条件を満たし、使用する色の数が最小であるような塗分けが
// 複数存在するときは、そのうちのどれを出力しても構わない

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Edge struct {
	ToNode int
	EdgeI  int
}

var stdin = func() *bufio.Scanner {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Split(bufio.ScanWords)
	return stdin
}()
var edges [][]Edge
var edgeColors []int
var maxColor int

func main() {
	n := readInt()

	// n木作成
	var root int
	edges = make([][]Edge, n+1)
	for i := 1; i < n; i++ {
		a := readInt()
		if i == 1 {
			root = a
		}
		b := readInt()
		edges[a] = append(edges[a], Edge{ToNode: b, EdgeI: i})
		edges[b] = append(edges[b], Edge{ToNode: a, EdgeI: i})
	}

	edgeColors = make([]int, n)
	depthFirstSearchForEdge(root, -1, -1)

	fmt.Println(maxColor - 1)

	for i := 1; i < n; i++ {
		fmt.Println(edgeColors[i])
	}
}

// readString() 読み込んだ文字列を返す
func readInt() int {
	stdin.Scan()
	x, err := strconv.Atoi(stdin.Text())
	if err != nil {
		panic(err)
	}
	return x
}

// depthFirstSearchForEdge() 深さ優先探索に似たものを用いて、辺に色を付ける
func depthFirstSearchForEdge(currentNode, fromNode, usedColor int) {
	color := 1
	for _, e := range edges[currentNode] {
		if e.ToNode == fromNode {
			continue
		}
		if color == usedColor {
			color++
			if color > maxColor {
				maxColor = color
			}
		}
		edgeColors[e.EdgeI] = color
		depthFirstSearchForEdge(e.ToNode, currentNode, color)
		color++
		if color > maxColor {
			maxColor = color
		}
	}
}
