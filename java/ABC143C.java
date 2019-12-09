package java;

import java.util.Scanner;
// -----問題文-----
// N匹のスライムが横一列に並んでいます。これらの色に関する情報が、長さNの英小文字
// から成る文字列Sで与えられます。左からi番目のスライムは、Sのi文字目に対応する
// 色を持っています。
// 同じ色を持ち隣接するスライムは融合し、色は変わらずに1匹のスライムとなります。
// この時、融合した後のスライムは、融合する前の各スライムが隣接していた他の
// スライムと隣接した状態になります。
// 最終的に存在するスライムは何匹となるでしょうか。
// -----制約-----
// 1≦N≦10^5
// |S|=N
// Sは英小文字から成る
// -----入力-----
// N
// S
// -----出力-----
// 最終的に存在するスライムの数を出力せよ。

public class ABC143C {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        int n = sc.nextInt();
        char[] slimesColor = sc.next().toCharArray();

        int slimeNum = 1;
        for (int i = 1; i < n; i++) {
            if (slimesColor[i] != slimesColor[i-1]) {
                slimeNum++;
            }
        }
        System.out.println(slimeNum);
    }
}