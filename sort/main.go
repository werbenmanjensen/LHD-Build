package main

import (
	"fmt"
	"sort"
)

func main() {
	list := []int{1,5,3}
	sort.Ints(list)
	fmt.Println(list)
}
