package main

import "fmt"

func Hello() string {
	return "Hello Go"
}

func printHelloBanner() {
	fmt.Println(Hello())
}

func main() {
	printHelloBanner()
}
