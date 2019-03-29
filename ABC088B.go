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
	sc.Split(bufio.ScanWords)
	
	n, alice, bob, count := nextInt(), 0, 0, 0
	var slice []int
	for i := 1; i <= n; i++ {
		slice = append(slice, nextInt())
	}
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	for i := 0; i < len(slice); i++ {
		count++
		if count % 2 != 0 {
			alice += slice[i]
		}else{
			bob += slice[i]
		}
	}
	fmt.Print(alice - bob)
}