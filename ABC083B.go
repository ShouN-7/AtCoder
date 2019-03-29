package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	allSum, limit, min, max := 0, nextInt(), nextInt(), nextInt()
	for i := 0; i <= limit; i++ {
		sum, s := 0, strconv.Itoa(i)
		if l := len(s); l > 1 {
			slice := s[0:l]
			for _, v := range slice {
				ci, _ := strconv.Atoi(string(v))
				sum += ci
			}
		}else{
			ci, _ := strconv.Atoi(s)
			sum = ci
		}
		if sum >= min && sum <= max {
			allSum += i
		}
	}
	fmt.Print(allSum)
}