package main

import "fmt"
import "os"

func main () {
    if len(os.Args) > 1 {
        var name = os.Args[1]
        fmt.Println("Hello", name)
    } else {
        fmt.Println("Usage: Supply name")
    }
}
