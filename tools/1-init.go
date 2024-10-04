package tools

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Initialization() []int {

	if len(os.Args) != 2 {
		fmt.Print("Missing arguments\n")
		os.Exit(0)
	}
	args := os.Args[1:]
	content, _ := os.ReadFile(args[0])
	numbers := strings.Fields(string(content))
	var arr []int
	for _, number := range numbers {
		n, _ := strconv.Atoi(number)
		arr = append(arr, n)
	}
	return arr
}
