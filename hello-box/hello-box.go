package main

import (
    "fmt"
    "os"
    "github.com/ozgio/strutil"
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
        output, _ = strutil.DrawCustomBox(output, 40, strutil.Center, strutil.SimpleBox9Slice(), "\n")
        fmt.Println(output)
    }

}

