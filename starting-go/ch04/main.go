package main

import "fmt"

func main() {
	s := make([]int, 5, 10)
	fmt.Println(s)
	s[1] = 3
	fmt.Println(s)

	fmt.Println(len(s))
	fmt.Println(cap(s))

	slice := []int{1, 2, 3}
	fmt.Println(len(slice))

	slice = append(slice, 5)
	fmt.Println(len(slice))

	m := make(map[int]string)
	m[1] = "US"
	m[81] = "Japan"
	m[86] = "China"

	// マップのリテラル
	m = map[int]string{1: "Taro", 2: "hanako", 3: "Jiro"}

	fmt.Println(m)
}
