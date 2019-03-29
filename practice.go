package main

import ( "fmt" )

func main() {
	var a, b, c int
	var s string
	fmt.Scan(&a)
	fmt.Scan(&b, &c)
	fmt.Scan(&s)

	str := fmt.Sprintf("%v %v\n", a + b + c, s)

	fmt.Print(str)

}