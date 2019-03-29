package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

const (
	ICHIMAN = 10000
	GOSEN   = 5000
	SEN     = 1000
)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	ss := strings.Split(nextLine(), " ")

	n, _ := strconv.Atoi(ss[0])
	y, _ := strconv.Atoi(ss[1])
	a, b, c := 0, 0, 0

	if y >= ICHIMAN {
		a = y / ICHIMAN
		for a >= 0 {
			if y-ICHIMAN*a == 0 {
				if n == a+b+c {
					break
				}
			}
			b = (y - ICHIMAN*a) / GOSEN
			for b >= 0 {
				if y == ICHIMAN*a+GOSEN*b {
					if n == a+b+c {
						break
					} else {
						b--
						continue
					}
				}
				c = (y - (ICHIMAN*a + GOSEN*b)) / SEN
				if y == ICHIMAN*a+GOSEN*b+SEN*c {
					if n == a+b+c {
						break
					}
				}
				b--
			}
			if n == a+b+c {
				break
			}
			a--
		}

	} else if y < ICHIMAN && y >= GOSEN {
		b = y / GOSEN
		for b >= 0 {
			if y != GOSEN*b {
				c = (y - GOSEN*b) / SEN
			}
			if y == GOSEN*b+SEN*c && n == a+b+c {
				break
			}
			b--
		}

	} else {
		c = y / SEN
	}

	if n == a+b+c && y == ICHIMAN*a+GOSEN*b+SEN*c {
		fmt.Printf("%v %v %v", a, b, c)
	} else {
		fmt.Print("-1 -1 -1")
	}
}
