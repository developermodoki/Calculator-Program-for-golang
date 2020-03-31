package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	fmt.Println("足し算だったら1、引き算だったら2と入力してください。:")
	var question int
	_, _ = fmt.Scan(&question)
	if question == 1 {
		fmt.Println("足される数を入力してください:")
		var question2 int
		_, _ = fmt.Scan(&question2)
		fmt.Println("足す数を入力してください:")
		var question3 int
		_, _ = fmt.Scan(&question3)
		fmt.Println(question2 + question3)
	}

	if question == 2 {
		fmt.Println("引かれる数を入力してください")
		var question4 int
		_, _ = fmt.Scan(&question4)
		fmt.Println("引く数を入力してください")
		var question5 int
		_, _ = fmt.Scan(&question5)
		fmt.Println(question4 - question5)
	}

}