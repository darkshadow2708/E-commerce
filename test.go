package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// greetUser prints a greeting with the user's name
func greetUser(name string) {
    fmt.Printf("Hello, %s! Welcome to Golang.\n", name)
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your name: ")
    
    input, _ := reader.ReadString('\n')
    name := strings.TrimSpace(input)

    greetUser(name)
}
