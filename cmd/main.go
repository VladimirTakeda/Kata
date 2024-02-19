package main

import (
	"Kata/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	expression, err := reader.ReadString('\n')
	if err != nil {
		panic("Can't read the expression")
	}
	fmt.Print(utils.Calculate(expression))
}
