package main

import (
    "fmt"
    "os"
    "github.com/ozgio/strutil"
)

func main () {

    value := ""
    if len(os.Args) > 1 {
        value = os.Args[1]
    } else {
        value = os.Getenv("HELLO_BOX_ARG")
    }
    if len(value) == 0 {
        value = "-h"
    }

    if value == "-h" || value == "--help" {
        fmt.Println("USAGE: Provide a string and we'll say hello! [can use env var HELLO_BOX_ARG]")
    } else {
        output := "Hello, " + value + "!"
        output, _ = strutil.DrawCustomBox(output, 40, strutil.Center, strutil.SimpleBox9Slice(), "\n")
        fmt.Println(output)
    }

    //For builds with CNB/Paketo:
    //    Using base run-image, two options (command overrides env var):
    //        docker run -e HELLO_BOX_ARG=world hello:base "hello-box sunshine"
    //    Using tiny run image, use:
    //        docker run -e HELLO_BOX_ARG=world hello:tiny


}

