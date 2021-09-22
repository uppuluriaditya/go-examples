package main

import (
	"fmt"

	esl "github.com/uppuluriaditya/go-examples/esl"
	"github.com/uppuluriaditya/go-examples/misc"
)

func main() {
	esl.StartMonitoring()
	fmt.Println(misc.FindSecondLargestNumber([]int{1, 4, 7, 5, 8}))
}
