package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"sort"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	n, dan, tmp := nextInt(), 0, 0
	var keis []int
	for i := 0; i < n; i++ {
		keis = append(keis, nextInt())
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keis)))
	for _, kei := range keis {
		if tmp == 0 {
			dan += 1
		}else{
			if (tmp > kei){
				dan += 1
			}
		}
		tmp = kei
	}
	fmt.Print(dan)
}