package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/XLXing/phonedata"
)

func main() {
	fmt.Println("请输入手机号:")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		pr, err := phonedata.Find(line)
		if err != nil {
			fmt.Printf("%s\n", err)
			fmt.Println("请重新输入手机号:")
		} else {
			fmt.Println(pr)
			fmt.Println("请输入手机号:")
		}
	}
	if err := input.Err(); err != nil {
		fmt.Println(os.Stderr)
		fmt.Printf("find phonedata: %v\n", err)
		os.Exit(1)
	}
}
