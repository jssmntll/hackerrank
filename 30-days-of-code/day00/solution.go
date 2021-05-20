package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()

    // Holds the string that was scanned
    inputString := scanner.Text()

    // Print the lines
    fmt.Println("Hello, World.")
    fmt.Println(inputString)
}
