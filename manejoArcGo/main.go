package main

import "fmt"

func printHello() {
	fmt.Println("Hello World")
}

func main() {
	defer printHello()
}

/* func main() {
	defer func() {
	   fmt.Println("Hello World")
	}()
  }*/
