package main

import "fmt"

func main() {
	var n int
	var x interface{}
	fmt.Println(n)         // => 初期値が0
	fmt.Printf("%#v\n", x) // => <nil> 具体的な値を持っていない状態を表す値

	q, r := div(19, 7)
	fmt.Printf("商=%d 剰余=%d\n", q, r)

	// 無名関数
	f := func(x, y int) int { return x + y }
	fmt.Println(f(1, 2))

	// 関数を返す関数
	returnFunc := returnFunc()
	returnFunc()

	// 関数を引数にとる関数
	callFunction(func() {
		fmt.Println("I'm a function")
	})

	// クロージャ
	late := later()
	fmt.Println(late("Golang"))
	fmt.Println(late("is"))

	// for
	for i := 0; i < 5; i++ {
		fmt.Println("hoge")
	}

	// 配列型とrange
	fruits := [3]string{"Apple", "Banana", "Cherry"}
	for i, s := range fruits {
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}

	// if
	if x, y := 1, 2; x < y {
		fmt.Printf("x(%d) is less then y(%d)\n", x, y)
	}

	// switch
	value := 3
	switch value {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("unknown")
	}

	// goto
	fmt.Println("A")
	goto L
	fmt.Println("B")

	L: /* ラベル */
		fmt.Println("C")

	// defer
	runDefer()
}

// int型のパラメーターa, bを受け取り、足し合わせた数値をint型で返す
func plus(x, y int) int {
	return x + y
}

func div(a, b int) (int, int) {
	q := a / b
	r := a % b

	return q, r
}

// 関数を返す関数
func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function.")
	}
}

// 関数を引数にとる関数
func callFunction(f func()) {
	f()
}

// クロージャー
func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}

}

// defer
func runDefer() {
	defer fmt.Println("defer")
	fmt.Println("done")
}

