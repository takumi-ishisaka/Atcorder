// -----問題文-----
// 長さNの値のわからない整数列Aがあります。長さN-1の整数列Bが与えられます。
// この時、Bi≧max（Ai, Ai+1）が成立することが分かっています。
// Aの要素の総和として考えらえる値の最大値を求めてください。
// -----制約-----
// ・入力は全て整数
// ・2≦N≦100
// ・0≦Bi≦10^5
// -----入力-----
// N
// B1 B2 .... BN-1
// -----出力-----
// Aの要素の総和として考えられる値の最大値を出力せよ。
#include<stdio.h>

// 配列情報格納に使用
int b[101], a[101];
// min() 小さいほうを返す
int min(int a, int b){
	return (a < b ? a : b);
}

int main() {
    int n;
    scanf("%d", &n);
	for(int i = 0; i < n-1; i++) {
        scanf("%d", &b[i]);
    }
    // Bi≧max(Ai, Ai+1)、Bi+1≧max(Ai+1, Ai+2)より、、
    // Ai+1≦min(Bi, Bi+1)となるから、Ai≦min(Bi-1, Bi)となる。
    // A1≦min(null, B1)、AN≦min(BN-1, null)で配列が完成する。
    // よって、要素の総和が最大となる配列を求めることができる。
    for(int i = 1; i < n - 1; i++) {
        a[i] = min(b[i - 1], b[i]);
    }
	a[0] = b[0]; 
    a[n - 1] = b[n - 2];

	int sum = 0;
	for(int i = 0; i < n; i++) sum += a[i];
	printf("%d\n", sum);
}