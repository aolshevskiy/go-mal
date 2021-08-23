package main

import (
	"bufio"
	"fmt"
	"os"
)

func READ(str string) string {
	return str
}

func EVAL(ast string, env string) string {
	return ast
}

func PRINT(exp string) string {
	return exp
}

func rep(str string) string {
	return PRINT(EVAL(READ(str), ""))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("user> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Println(rep(text))
	}
}
