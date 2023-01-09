package main

import (
	// 	出力用パッケージ
	"fmt"
	// 数学系パッケージ
	"math"
	// 複素数パッケージ
	"math/cmplx"
	// ランダム数値取得用

	// 時間取得のパッケージ
	"time"
)

// 変数定義
// varを使い変数を宣言できる。まとめて宣言して末尾で型を宣言できる。
// constを使い定数を宣言できる。character,string,numeric,booleanのみで利用可
// 初期値を与えなかった場合、bool→false, intなど数値→0, string→""
var c, java, python bool

// 初期値を入れた場合、型の記述を省略できる。
var i int = 1
var j = 2

// 関数の中では:=を使って省略形で変数の宣言ができる。
// 定数は:=を使って宣言できない
func printK() {
	k := 3
	fmt.Println("K is", k)
}

func printT() {
	var n int64 = 1 << 34
	var cmpl complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Printf("Type: %T, Value: %v\n", n, n)
	fmt.Printf("Type: %T, Value: %v\n", cmpl, cmpl)
}

// 関数定義
// func 関数名(引数 引数の型) 返り値の型 {}
// 二つの引数の型が同じ場合(x,y int)のように宣言可#swap
// 返り値にx,yと名前を付け、returnに何も記述せずに帰すことができる#split
func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 繰り返し処理
// for 初期化;条件式;後処理 {}
// 初期化と後処理の記述は任意
func sum(a, d, n int) int {
	sum := a
	for i := 1; i <= n; i++ {
		sum += (i * d)
	}
	return sum
}

// goにはWhileは無く、forのみで表現する。
func while(limit int) int {
	sum := 1
	for sum < limit {
		sum += sum
	}
	return sum
}

// 分岐処理
// if:forと同様に条件式の前に簡単なステートメントを記述することができる
func pow(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	}
	return limit
}

// switch: 条件ステートメントを書けばその条件に一致するcaseブロックへ、
// 省略するとcase文の条件式がtrueになるブロックへ分岐する。
func greet() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}
}

// 遅延評価
// defer: 元の呼び出し関数がreturnするまで、渡した関数を評価しない
// deferへ渡した関数が複数ある場合はLIFOで評価される
func delay() {
	defer fmt.Println("後に出力されます")
	fmt.Println("先に出力されます")
	for i := 1; i < 6; i++ {
		defer fmt.Printf("%v番目に呼び出し\n", i)
	}
}

// ポインタ
// 変数Tのポインタは*T型
// p=&xでxにアクセスするポインタをpに渡せる
// *pでポインタを通してxにアクセスできる
// 構造体
// 構造名 struct {フィールド名 型}で構造体を定義できる。
// 構造体の一部だけを定義して初期化すると、他の値は自動的に初期値が入る。
func point(p1 int) {
	x := &p1
	fmt.Println(x, *x)
	type Vertex struct {
		X int
		Y int
	}
	v := Vertex{X: 1}
	p2 := &v
	p2.X = 10
	fmt.Println(v, p2, *p2)
}

// 配列
// Array: [n]TはT型のn個の変数の配列を表す。
// []T型はスライスを示す。スライスは配列の参照なので、配列自体は保持しない。
// sliceは長さ（length）と容量（capacity）を持っている
// それぞれlen()とcap()で取得できる。
// sliceの長さ、容量がゼロの時はnilになる
func printL() {
	var l [10]int
	s := l[1:5]
	s[2] = 2
	fmt.Printf("Type: %T Value: %v\n", l, l)
	fmt.Printf("Type: %T Value: %v\n", s, s)
	p := []int{1, 2, 3, 4, 5}
	q := []struct {
		x int
		y bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
	}
	fmt.Printf("Type: %T Value: %v\n", p, p)
	fmt.Printf("Type: %T Value: %v\n", q, q)
	var n []int
	if n == nil {
		fmt.Println("n is nil!")
	}
}

func main() {
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println("hello world!")
	// fmt.Println("My favorite number is", rand.Intn(100))
	// fmt.Printf("Pi is %g\n", math.Pi)
	// fmt.Println(add(10, 10))
	// pineapple, apple := swap("pineapple", "apple")
	// fmt.Println(pineapple, apple)
	// fmt.Println(split(35))
	// fmt.Println(i, j, c, java, python)
	// printK()
	// printT()
	// fmt.Println(sum(3, 2, 10))
	// fmt.Println(while(100))
	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 20),
	// )
	// greet()
	// delay()
	// point(29)
	printL()
}
