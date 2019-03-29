package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil{
		panic(e)
	}

	return i
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()

	var slice []int
	for i := 0; i < n; i++ {
		slice = append(slice, nextInt())
	}

	b := 0
	flg := true
	for flg {
		for i, v := range slice{
			if (v % 2 != 0){
				flg = false
				break
			}else{
				slice[i] = v / 2
			}
		}
		if (flg){ 
			b += 1
		}
	}

	fmt.Print(b)
}