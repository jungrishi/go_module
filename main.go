package main

import (
 	"fmt"
 	"github.com/jungrishi/go_module/calc"
 )
 
 func main() {
 	arr := [5]int{1,2,3,4,5}
  sum := calc.Add(arr[:]...)
 	fmt.Println(sum)
 }
