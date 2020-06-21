package main

import "fmt"

func main() {
	// 変数宣言
	// var型の宣言は関数外でも宣言が可能
	// var i int = 1
	// var f64 float64 = 1.2
	// var s string = "test"
	// var t bool = true
	// var f bool = false
	var (
		i   int     = 1
		f64 float64 = 1.2
		s   string  = "test"
		t   bool    = true
		f   bool    = false
	)
	// var (
	// 	i   int
	// 	f64 float64
	// 	s   string
	// 	t   bool
	// 	f   bool
	// )
	fmt.Println(i, f64, s, t, f)

	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	// 下記は変数の型を出力
	fmt.Printf("%T\n", xf64)
	fmt.Printf("%T\n", xi)
}
