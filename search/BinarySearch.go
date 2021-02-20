package search

import (
	"bufio"
	"fmt"
	"learn_go/algorithm"
	"os"
	"sort"
	"strconv"
)

func BinarySearch(key int, a []int) int {
	lo := 0
	hi := len(a) - 1
	for {
		if lo > hi {
			break
		}
		mid := lo + (hi - lo) / 2
		if key < a[mid] {
			hi = mid - 1
		} else if key > a[mid] {
			lo = mid + 1
		}else{
			return mid
		}
	}
	return -1
}

func Use()  {
	scanner := bufio.NewScanner(os.Stdin)
	var whiteList []int
	for scanner.Scan(){
		if s, err := strconv.Atoi(scanner.Text()); err == nil{
			whiteList = append(whiteList, s)
		}
	}
	sort.Ints(whiteList)
	fmt.Println(algorithm.BinarySearch(30, whiteList))

}

