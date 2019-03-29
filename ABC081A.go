package main

import( "fmt" )

func main() {
	b := 0
	var s string
	fmt.Scan(&s)

	var x, y = 0, 1
	for i := 0; i < 3; i++{
		 temp := s[ x : y ]

		 if (temp == "1"){
			 b += 1
		 }
		 x, y = y, y + 1
	}
	
	fmt.Print(b)
}