package main

import "log"

func soma(a int, b int) int {
	return a + b
}

func main() {
	log.Println(soma(5, 5))
}
