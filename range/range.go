package main

import "fmt"

func main() {
	// よくあるやつ
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	// goのrangemを使用してやる
	for i, v := range l {
		fmt.Println(i, v)
	}

	// goのrangemを使用してやる キーを出力しない場合
	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for _, v := range m {
		fmt.Println(v)
	}
}
