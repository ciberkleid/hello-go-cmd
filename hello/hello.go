package main

import (
    "fmt"
    "os"
)

func main () {

    value := "-h"
    if len(os.Args) > 1 {
        value = os.Args[1]
    }

    if value == "-h" || value == "--help" {
        fmt.Println("USAGE: Provide a string and we'll say hello!")
    } else {
        output := "Hello, " + value + "!"
        fmt.Println(output)
    }

}

