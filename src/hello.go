package main

import (
    "fmt"
    "github.com/ozgio/strutil"
    "os"
    "strconv"
    "time"
)

func main () {

    value := ""
    if len(os.Args) > 1 {
        value = os.Args[1]
    } else {
        value = os.Getenv("HELLO_ARG")
    }
    if len(value) == 0 {
        value = "-h"
    }

    if value == "-h" || value == "--help" {
        fmt.Println("USAGE: Provide a string and we'll say hello!")
    } else {
        output := "Hello, " + value + "!"
        output, _ = strutil.DrawCustomBox(output, 40, strutil.Center, strutil.SimpleBox9Slice(), "\n")
        fmt.Println(output)
    }

    IsSleepEnabled, _ := strconv.ParseBool(os.Getenv("HELLO_SLEEP"))
    if IsSleepEnabled {
        fmt.Printf("Sleeping 30 seconds...\n")
        time.Sleep(30 * time.Second)
    }

}

//For builds with CNB/Paketo:
//    Using base run-image, two options (command overrides env var):
//        docker run -e HELLO_ARG=world hello:base "hello sunshine"
//    Using tiny run image, use:
//        docker run -e HELLO_BOX_ARG=world hello:tiny

