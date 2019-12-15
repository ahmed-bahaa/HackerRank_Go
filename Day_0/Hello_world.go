package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUTo.
	reader := bufio.NewReader(os.Stdin)
	inputStr, _ := reader.ReadString('\n')
	inputStr = strings.TrimSpace(inputStr)
	fmt.Print("Hello, World. \n")
	fmt.Println(inputStr)
}
