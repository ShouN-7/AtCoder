package main

import(
	"fmt"
)

func main() {
	var x, a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&x)

	p := 0

	switch {
	case x >= 500:
		n := x / 500
		if n > a {
			n = a
		}
		for n >= 0 {
			y := x - 500 * n
			if y == 0 {
				p += 1
			}else{
				p += calcBAndC(y, b, c)
			}
			n--
		}
	case x > 50 && x < 500:
		p += calcBAndC(x, b, c)
	default:
		p += 1
	}
	fmt.Print(p)
}

func calcBAndC(x, b, c int) int{
	p, n := 0, x / 100
	if n > b {
		n = b
	}
	for n >= 0 {
		if y := x - 100 * n; y == 0{
			p += 1
		}else{
			if m := y / 50; m > c {
				break
			}else{
				p += 1
			}
		}
		n--
	}
	return p
}