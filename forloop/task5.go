package main

import "fmt"

func main() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)
	for i := 2; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(c)
			break
		}
	}
}