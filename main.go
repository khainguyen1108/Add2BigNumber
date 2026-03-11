package main

import (
	add2num "ADD2NUM/add2num"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// tao file log
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("[main.file.err]", err)
		return
	}

	defer file.Close()

	log.SetOutput(file)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Number1: ")
	a, _ := reader.ReadString('\n')

	fmt.Print("Number2: ")
	b, _ := reader.ReadString('\n')

	a = strings.TrimSpace(a)
	b = strings.TrimSpace(b)

	myBigNumber := add2num.MyBigNumber{}
	result := myBigNumber.Sum(a, b)

	fmt.Println("Result:", result)

	log.Println("Input 1:", a, "Input 2:", b, "Result:", result)
}
