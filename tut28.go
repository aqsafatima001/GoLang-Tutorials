package main

import "fmt"

// Declare variable activeUserCount
var activeUserCount int = 0

func entry() {
    // Hint: you can use the "++" operator to increment a variable by 1
    activeUserCount++
}

func exit() {
    // Hint: you can use the "--" operator to decrement a variable by 1
    activeUserCount--
}

func main() {
    entry()
    entry()
    exit()
    exit()
    entry()
    entry()
    fmt.Println(activeUserCount)
}